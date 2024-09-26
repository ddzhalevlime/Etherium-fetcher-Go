package web3

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"os"
	"time"

	"eth-fetcher.ddzhalev.net/internal/models"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PersonInfoContractInteractor struct {
	httpClient *ethclient.Client
	wsClient   *ethclient.Client
	contract   *SimplePersonInfoContract
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    *big.Int
}

func NewPersonInfoContractInteractor() (*PersonInfoContractInteractor, error) {
	httpClient, err := ethclient.Dial(os.Getenv("ETH_NODE_URL"))
	if err != nil {
		return nil, err
	}

	wsClient, err := ethclient.Dial(os.Getenv("ETH_SOCKET_URL"))
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	contract, err := newContractInstance(httpClient)
	if err != nil {
		return nil, err
	}

	chainID, err := httpClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &PersonInfoContractInteractor{
		httpClient: httpClient,
		wsClient:   wsClient,
		contract:   contract,
		privateKey: privateKey,
		address:    address,
		chainID:    chainID,
	}, nil
}

func (pci *PersonInfoContractInteractor) SetPersonInfo(name string, age int) (string, bool, error) {
	auth, err := pci.getTransactOpts()
	if err != nil {
		return "", false, err
	}

	tx, err := pci.contract.SetPersonInfo(auth, name, big.NewInt(int64(age)))
	if err != nil {
		return "", false, err
	}

	receipt, err := pci.waitForTx(tx.Hash())
	if err != nil {
		return tx.Hash().Hex(), false, err
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		return tx.Hash().Hex(), true, nil
	} else {
		return tx.Hash().Hex(), false, nil
	}
}

func (pci *PersonInfoContractInteractor) GetPersonInfo(index int) (string, int, error) {
	name, age, err := pci.contract.GetPersonInfo(&bind.CallOpts{}, big.NewInt(int64(index)))
	if err != nil {
		return "", 0, err
	}
	return name, int(age.Int64()), nil
}

func (pci *PersonInfoContractInteractor) GetPersonsCount() (int, error) {
	count, err := pci.contract.GetPersonsCount(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return int(count.Int64()), nil
}

func (pci *PersonInfoContractInteractor) getTransactOpts() (*bind.TransactOpts, error) {
	nonce, err := pci.httpClient.PendingNonceAt(context.Background(), pci.address)
	if err != nil {
		return nil, err
	}

	gasPrice, err := pci.httpClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pci.privateKey, pci.chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

func (pci *PersonInfoContractInteractor) ListenForEvents(ctx context.Context, eventModel *models.PersonInfoEventModel) {
	wsContract, err := newContractInstance(pci.wsClient)
	if err != nil {
		log.Printf("Failed to create WebSocket contract instance: %v", err)
		return
	}

	filterer := &SimplePersonInfoContractFilterer{contract: wsContract.SimplePersonInfoContractFilterer.contract}

	sink := make(chan *SimplePersonInfoContractPersonInfoUpdated)
	sub, err := filterer.WatchPersonInfoUpdated(&bind.WatchOpts{Context: ctx}, sink, nil)
	if err != nil {
		log.Printf("Failed to watch for PersonInfoUpdated events: %v", err)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			err := eventModel.Insert(&models.PersonInfoEvent{
				PersonIndex:     int(event.PersonIndex.Int64()),
				PersonName:      event.NewName,
				PersonAge:       int(event.NewAge.Int64()),
				TransactionHash: event.Raw.TxHash.Hex(),
			})
			if err != nil {
				log.Printf("Failed to insert event: %v", err)
			} else {
				log.Printf("Inserted new PersonInfoUpdated event: %s", event.Raw.TxHash.Hex())
			}
		case err := <-sub.Err():
			log.Printf("Event subscription error: %v", err)
			return
		case <-ctx.Done():
			log.Println("Event listener stopped")
			return
		}
	}
}

func newContractInstance(client *ethclient.Client) (*SimplePersonInfoContract, error) {
	return NewSimplePersonInfoContract(common.HexToAddress(os.Getenv("SIMPLE_PERSON_INFO_CONTRACT_ADDRESS")), client)
}

func (pci *PersonInfoContractInteractor) waitForTx(txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	for {
		receipt, err := pci.httpClient.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		if err != ethereum.NotFound {
			return nil, err
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second):
		}
	}
}

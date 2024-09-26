package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"eth-fetcher.ddzhalev.net/internal/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang-jwt/jwt"
)

func handleTransactionHashesQueryString(r *http.Request) ([]string, error) {
	transactionHashes := r.URL.Query()["transactionHashes"]
	if len(transactionHashes) == 0 {
		return nil, fmt.Errorf("missing transactionHashes parameter")
	}

	var flattenedHashes []string
	for _, hash := range transactionHashes {
		flattenedHashes = append(flattenedHashes, strings.Split(hash, ",")...)
	}

	uniqueHashes := make(map[string]bool)
	var result []string
	for _, hash := range flattenedHashes {
		if hash = strings.TrimSpace(hash); hash != "" {
			if !uniqueHashes[hash] {
				uniqueHashes[hash] = true
				result = append(result, hash)
			}
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no valid transaction hashes provided")
	}

	return result, nil
}

func mapTransactionToModel(ethTx *types.Transaction, receipt *types.Receipt, blockHeader *types.Header) (*models.Transaction, error) {
	value := ethTx.Value().String()

	fromAddress, err := types.Sender(types.NewLondonSigner(ethTx.ChainId()), ethTx)
	if err != nil {
		return nil, fmt.Errorf("failed to extract sender address: %w", err)
	}

	var contractAddress string
	if receipt.ContractAddress != (common.Address{}) {
		contractAddress = receipt.ContractAddress.Hex()
	}

	return &models.Transaction{
		TransactionHash:   ethTx.Hash().Hex(),
		TransactionStatus: int(receipt.Status),
		BlockHash:         blockHeader.Hash().Hex(),
		BlockNumber:       blockHeader.Number.Uint64(),
		From:              fromAddress.Hex(),
		To:                ethTx.To().Hex(),
		ContractAddress:   contractAddress,
		LogsCount:         len(receipt.Logs),
		Input:             common.Bytes2Hex(ethTx.Data()),
		Value:             value,
	}, nil
}

func (app *application) responseJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		app.serverError(w, r, fmt.Errorf("failed to encode response: %w", err))
	}
}

func (app *application) fetchTransactions(ctx context.Context, hashStrings []string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	for _, hashString := range hashStrings {
		tx, err := app.fetchAndStoreTransaction(ctx, hashString)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (app *application) fetchAndStoreTransaction(ctx context.Context, hashString string) (*models.Transaction, error) {
	tx, err := app.transactions.Get(hashString)
	if err == nil {
		return tx, nil
	}

	if err.Error() != "transaction not found" {
		return nil, fmt.Errorf("failed to get transaction %s from the DB: %w", hashString, err)
	}

	ethTx, isPending, err := app.ethClient.TransactionByHash(ctx, common.HexToHash(hashString))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction %s: %w", hashString, err)
	}

	if isPending {
		return nil, fmt.Errorf("transaction %s is still pending", hashString)
	}

	receipt, err := app.ethClient.TransactionReceipt(ctx, ethTx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch receipt for transaction %s: %w", hashString, err)
	}

	blockHeader, err := app.ethClient.HeaderByHash(ctx, receipt.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch block header for transaction %s: %w", hashString, err)
	}

	tx, err = mapTransactionToModel(ethTx, receipt, blockHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to convert transaction %s: %w", hashString, err)
	}

	err = app.transactions.Insert(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to insert transaction %s: %w", hashString, err)
	}

	return tx, nil
}

func extractTransactionIds(transactions []*models.Transaction) []int {
	ids := make([]int, len(transactions))
	for i, tx := range transactions {
		ids[i] = tx.ID
	}
	return ids
}

func hashWithJwtSecret(data string) string {
	jwtSecret := os.Getenv("JWT_SECRET")
	dataHash := hmac.New(sha256.New, []byte(jwtSecret))
	dataHash.Write([]byte(data))
	return hex.EncodeToString(dataHash.Sum(nil))
}

func (app *application) parseToken(tokenString string) (*jwt.Token, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is not set in the environment")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return token, nil
}

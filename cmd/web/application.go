package main

import (
	"log/slog"

	"eth-fetcher.ddzhalev.net/internal/models"
	"eth-fetcher.ddzhalev.net/internal/web3"
	"github.com/ethereum/go-ethereum/ethclient"
)

type application struct {
	logger             *slog.Logger
	transactions       *models.TransactionModel
	users              *models.UserModel
	personInfoEvents   *models.PersonInfoEventModel
	ethClient          *ethclient.Client
	jwtSecret          string
	contractInteractor *web3.PersonInfoContractInteractor
}

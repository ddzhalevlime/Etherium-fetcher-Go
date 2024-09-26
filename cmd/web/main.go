package main

import (
	"database/sql"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	"eth-fetcher.ddzhalev.net/internal/models"
	"eth-fetcher.ddzhalev.net/internal/web3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	loadEnvFile()

	addr := flag.String("addr", os.Getenv("API_PORT"), "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("DB_CONNECTION_URL"), "PostgreSQL data source name")
	ethNodeURL := flag.String("ethnode", os.Getenv("ETH_NODE_URL"), "Ethereum node URL")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error("failed to open database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	ethClient, err := ethclient.Dial(*ethNodeURL)
	if err != nil {
		logger.Error("failed to connect to Ethereum node", "error", err)
		os.Exit(1)
	}

	contractInteractor, err := web3.NewPersonInfoContractInteractor()
	if err != nil {
		log.Fatalf("Failed to create contract interactor: %v", err)
	}

	app := &application{
		logger:             logger,
		transactions:       &models.TransactionModel{DB: db},
		users:              &models.UserModel{DB: db},
		personInfoEvents:   &models.PersonInfoEventModel{DB: db},
		ethClient:          ethClient,
		jwtSecret:          os.Getenv("JWT_SECRET"),
		contractInteractor: contractInteractor,
	}

	app.StartEventListener()

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", slog.String("addr", *addr))

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	userModel := &models.UserModel{DB: db}
	transactionModel := &models.TransactionModel{DB: db}
	personModel := &models.PersonInfoEventModel{DB: db}

	if err := userModel.CreateTable(); err != nil {
		return nil, err
	}

	if err := transactionModel.CreateTable(); err != nil {
		return nil, err
	}

	if err := personModel.CreateTable(); err != nil {
		return nil, err
	}

	if err := userModel.InitializeDefaultUsers(hashWithJwtSecret); err != nil {
		return nil, err
	}

	return db, nil
}

func loadEnvFile() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

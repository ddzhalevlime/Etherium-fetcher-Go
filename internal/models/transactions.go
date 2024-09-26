package models

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Transaction struct {
	ID                int    `json:"id"`
	TransactionHash   string `json:"transactionHash"`
	TransactionStatus int    `json:"transactionStatus"`
	BlockHash         string `json:"blockHash"`
	BlockNumber       uint64 `json:"blockNumber"`
	From              string `json:"from"`
	To                string `json:"to"`
	ContractAddress   string `json:"contractAddress"`
	LogsCount         int    `json:"logsCount"`
	Input             string `json:"input"`
	Value             string `json:"value"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (m *TransactionModel) CreateTable() error {
	_, err := m.DB.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			transactionHash VARCHAR(66) UNIQUE NOT NULL,
			transactionStatus INTEGER,
			blockHash VARCHAR(66),
			blockNumber BIGINT,
			fromAddress VARCHAR(42),
			toAddress VARCHAR(42),
			contractAddress VARCHAR(42),
			logsCount INTEGER,
			input TEXT,
			value TEXT
		)
	`)
	return err
}

func (m *TransactionModel) Insert(tx *Transaction) error {
	query := `
        INSERT INTO transactions (transactionHash, transactionStatus, blockHash, blockNumber, fromAddress, toAddress, contractAddress, logsCount, input, value)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `
	_, err := m.DB.Exec(query, tx.TransactionHash, tx.TransactionStatus, tx.BlockHash, tx.BlockNumber, tx.From, tx.To, tx.ContractAddress, tx.LogsCount, tx.Input, tx.Value)
	return err
}

func (m *TransactionModel) Get(hash string) (*Transaction, error) {
	query := `
		SELECT *        
		FROM transactions
        WHERE transactionHash = $1
    `
	tx := &Transaction{}
	err := m.DB.QueryRow(query, hash).Scan(&tx.ID, &tx.TransactionHash, &tx.TransactionStatus, &tx.BlockHash, &tx.BlockNumber, &tx.From, &tx.To, &tx.ContractAddress, &tx.LogsCount, &tx.Input, &tx.Value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}
	return tx, nil
}

func (m *TransactionModel) GetMultipleByIDs(ids []int) ([]*Transaction, error) {
	query := `
        SELECT *
        FROM transactions
        WHERE id = ANY($1)
    `
	return m.getMultipleTransactions(query, pq.Array(ids))
}

func (m *TransactionModel) GetAll() ([]*Transaction, error) {
	query := `
		SELECT *
		FROM transactions
		ORDER BY id
	`
	return m.getMultipleTransactions(query, nil)
}

func (m *TransactionModel) getMultipleTransactions(query string, param interface{}) ([]*Transaction, error) {
	var rows *sql.Rows
	var err error

	if param != nil {
		rows, err = m.DB.Query(query, param)
	} else {
		rows, err = m.DB.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return m.scanTransactions(rows)
}

func (m *TransactionModel) scanTransactions(rows *sql.Rows) ([]*Transaction, error) {
	transactions := []*Transaction{}
	for rows.Next() {
		tx := &Transaction{}
		err := rows.Scan(
			&tx.ID,
			&tx.TransactionHash,
			&tx.TransactionStatus,
			&tx.BlockHash,
			&tx.BlockNumber,
			&tx.From,
			&tx.To,
			&tx.ContractAddress,
			&tx.LogsCount,
			&tx.Input,
			&tx.Value,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

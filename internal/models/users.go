package models

import (
	"database/sql"

	"github.com/lib/pq"
)

type User struct {
	ID                     int    `json:"id"`
	Username               string `json:"username"`
	PasswordHash           string `json:"-"`
	SearchedTransactionIds []int  `json:"searchedTransactionIds"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) CreateTable() error {
	_, err := m.DB.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(50) UNIQUE NOT NULL,
				passwordHash TEXT,
				searchedTransactionIds INT[]
			)
		  `)

	return err
}

func (m *UserModel) InitializeDefaultUsers(hashFunc func(string) string) error {
	defaultUsers := map[string]string{
		"alice": "alice",
		"bob":   "bob",
		"carol": "carol",
		"dave":  "dave",
	}

	for username, password := range defaultUsers {
		if err := m.InsertIfNotExists(username, password, hashFunc); err != nil {
			return err
		}
	}
	return nil
}

func (m *UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users (username, passwordHash)
		VALUES ($1, $2)
	`
	_, err := m.DB.Exec(query, user.Username, user.PasswordHash)
	return err
}

func (m *UserModel) InsertIfNotExists(username, password string, hashFunc func(string) string) error {
	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		passwordHash := hashFunc(password)
		_, err = m.DB.Exec("INSERT INTO users (username, passwordHash) VALUES ($1, $2)", username, passwordHash)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *UserModel) InsertTransactionIds(username string, newTransactionIds []int) error {
	query := `
		UPDATE users
		SET searchedTransactionIds = ARRAY(
			SELECT DISTINCT unnest(ARRAY_CAT(searchedTransactionIds, $1::int[]))
		)
		WHERE username = $2
	`
	_, err := m.DB.Exec(query, pq.Array(newTransactionIds), username)
	return err
}

func (m *UserModel) Get(username string) (*User, error) {
	query := `
		SELECT *
		FROM users
		WHERE username = $1
	`
	user := &User{}
	var searchedTransactionIds pq.Int32Array

	err := m.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &searchedTransactionIds)
	if err != nil {
		return nil, err
	}

	user.SearchedTransactionIds = make([]int, len(searchedTransactionIds))
	for i, v := range searchedTransactionIds {
		user.SearchedTransactionIds[i] = int(v)
	}

	return user, nil
}

package models

import (
	"database/sql"
)

type PersonInfoEvent struct {
	ID              int    `json:"id"`
	PersonIndex     int    `json:"personIndex"`
	PersonName      string `json:"personName"`
	PersonAge       int    `json:"personAge"`
	TransactionHash string `json:"TransactionHash"`
}

type PersonInfoEventModel struct {
	DB *sql.DB
}

func (m *PersonInfoEventModel) CreateTable() error {
	_, err := m.DB.Exec(`
		CREATE TABLE IF NOT EXISTS personInfoEvents (
			id SERIAL PRIMARY KEY,
			personIndex INTEGER UNIQUE NOT NULL,
			personName TEXT NOT NULL,
			personAge INTEGER NOT NULL,
			transactionHash TEXT NOT NULL
		)
	`)

	return err
}

func (m *PersonInfoEventModel) Insert(event *PersonInfoEvent) error {
	query := `
		INSERT INTO personInfoEvents (PersonIndex, PersonName, PersonAge, TransactionHash)
		VALUES ($1, $2, $3, $4)
	`
	_, err := m.DB.Exec(query, event.PersonIndex, event.PersonName, event.PersonAge, event.TransactionHash)
	return err
}

func (m *PersonInfoEventModel) GetAll() ([]*PersonInfoEvent, error) {
	query := `
		SELECT *
		FROM personInfoEvents
	`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*PersonInfoEvent
	for rows.Next() {
		event := &PersonInfoEvent{}
		err := rows.Scan(&event.ID, &event.PersonIndex, &event.PersonName, &event.PersonAge, &event.TransactionHash)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

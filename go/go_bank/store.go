package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) (*Account, error)
	GetAccount(int) (*Account, error)
	UpdateAccount(*Account) error
	DeleteAccount(int) error
}

type PostgresStore struct {
	db *sql.DB
}

/**
 * currently requires you to run: `docker run --name go_bank_db -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=go_bank_db -p 5432:5432 -d postgres`
 */
func NewPostgresStore() (*PostgresStore, error) {
	// connStr := "user=postgres dbname=go_bank_db password=postgres sslmode=verify-full"
	connStr := "user=postgres dbname=go_bank_db password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {

	return s.createAccountTable()
}

// region DB Migrations
func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account
    (
    id SERIAL PRIMARY KEY,
    first_name BPCHAR NOT NULL,
    last_name BPCHAR NOT NULL,
    account_number UUID UNIQUE,
    balance DECIMAL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    )`

	_, err := s.db.Exec(query)

	return err
}

// endregion
func (s *PostgresStore) CreateAccount(acc *Account) (*Account, error) {
	query := `
    INSERT INTO account (
        first_name,
        last_name,
        account_number,
        balance,
    ) VALUES ($1, $2, $3, $4)
    RETURNING id, first_name, last_name, account_number, balance, created_at, updated_at
    `
	row, err := s.db.Exec(query, acc.FirstName, acc.LastName, acc.AccountNumber, acc.Balance)

	if err != nil {
		return nil, err
	}
	fmt.Print(row)
	return nil, nil
}
func (s *PostgresStore) GetAccount(int) (*Account, error) {
	return nil, nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(int) error {
	return nil
}

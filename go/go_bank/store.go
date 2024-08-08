package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*CreateAccountAccount) (*Account, error)
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
// pg helpers region
func scanRowIntoAccount(row *sql.Rows) (*Account, error) {
	account := new(Account)
	// must be in the same order as the columns in the database
	err := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return account, err
}

// endregion
func (s *PostgresStore) CreateAccount(acc *CreateAccountAccount) (*Account, error) {
	query := `
    INSERT INTO account (
        first_name,
        last_name,
        account_number,
        balance
    ) VALUES ($1, $2, $3, $4)
    RETURNING id, first_name, last_name, account_number, balance, created_at, updated_at
    `
	rows, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.AccountNumber, acc.Balance)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanRowIntoAccount(rows)
	}

	return nil, fmt.Errorf("Failed to create an account")
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

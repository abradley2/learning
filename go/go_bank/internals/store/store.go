package store

import (
	"database/sql"
	t "github.com/wley3337/learning/tree/main/go/go_bank/internals/types"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*t.CreateAccountAccount) (*t.Account, error)
	GetAccountByID(int) (*t.Account, error)
	GetAccounts() ([]*t.Account, error)
	UpdateAccount(*t.Account) error
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

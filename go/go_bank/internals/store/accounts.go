package store

import (
	t "github.com/wley3337/learning/tree/main/go/go_bank/types"

	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

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
func scanRowIntoAccount(row *sql.Rows) (*t.Account, error) {
	account := new(t.Account)
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
func (s *PostgresStore) CreateAccount(acc *t.CreateAccountAccount) (*t.Account, error) {
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

func (s *PostgresStore) GetAccounts() ([]*t.Account, error) {
	query := `
        SELECT
            id,
            first_name,
            last_name,
            account_number,
            balance,
            created_at,
            updated_at
        FROM account
    `
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	accounts := []*t.Account{}
	for rows.Next() {
		account, err := scanRowIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *PostgresStore) GetAccountByID(accountID int) (*t.Account, error) {
	query := `
        SELECT
            id,
            first_name,
            last_name,
            account_number,
            balance,
            created_at,
            updated_at
        FROM account
        WHERE id=$1
    `

	rows, err := s.db.Query(query, accountID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanRowIntoAccount(rows)
	}
	return nil, nil
}
func (s *PostgresStore) UpdateAccount(*t.Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(accountID int) error {
	query := `DELETE from account where id=$1`

	_, err := s.db.Query(query, accountID)

	if err != nil {
		return err
	}

	return nil
}

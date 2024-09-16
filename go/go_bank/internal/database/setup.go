package database

import "database/sql"

func Setup(db *sql.DB) error {
	return createAccountTable(db)
}

func createAccountTable(db *sql.DB) error {
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

	_, err := db.Exec(query)

	return err
}

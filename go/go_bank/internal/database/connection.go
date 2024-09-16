package database

import "database/sql"

func NewConnection() (*sql.DB, error) {
	// connStr := "user=postgres dbname=go_bank_db password=postgres sslmode=verify-full"
	connStr := "user=postgres dbname=go_bank_db password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

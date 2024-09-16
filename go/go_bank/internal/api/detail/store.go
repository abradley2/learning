package detail

import (
	t "github.com/wley3337/learning/tree/main/go/go_bank/types"

	"database/sql"

	_ "github.com/lib/pq"
)

type Store interface {
	GetAccountByID(int) (t.Account, error)
	DeleteAccount(int) error
}

type defaultStore struct {
	db *sql.DB
}

func (s defaultStore) GetAccountByID(accountID int) (t.Account, error) {
	acct := t.Account{}
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
		return acct, err
	}

	for rows.Next() {
		err = acct.FromRow(rows)
		return acct, err
	}
	return acct, nil
}

func (s defaultStore) DeleteAccount(accountID int) error {
	query := `DELETE from account where id=$1`

	_, err := s.db.Query(query, accountID)

	if err != nil {
		return err
	}

	return nil
}

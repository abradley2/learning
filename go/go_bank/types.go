package main

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountAccount struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber uuid.UUID `json:"accountNumber"`
	Balance       int64     `json:"balance"`
}

// in order to user scanRowIntoAccount, the order of the
// struct must match the column order in the db
type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber uuid.UUID `json:"accountNumber"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: uuid.New(),
	}
}

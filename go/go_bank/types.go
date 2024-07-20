package main

import (
	"math/rand"

	"github.com/google/uuid"
)

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber uuid.UUID `json:"accountNumber"`
	Balance       int64     `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:            rand.Intn(10000),
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: uuid.New(),
	}
}

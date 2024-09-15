package api

import (
	t "github.com/wley3337/learning/tree/main/go/go_bank/types"

	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch method := r.Method; method {
	case "GET":
		return s.handleGetAccounts(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	default:
		return fmt.Errorf("Method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		log.Println("Error getting accounts")
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

type CreateAccountRequestBody struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber uuid.UUID `json:"accountNumber"`
	Balance       int64     `json:"balance"`
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateAccountRequestBody)

	// this needs to parse the JSON from the body
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	lastFour := req.AccountNumber.String()[len(req.AccountNumber.String())-4:]

	log.Println("Creating account from:",
		req.FirstName,
		req.LastName,
		lastFour)

	// created a struct that the create account accepts
	account := new(t.CreateAccountAccount)
	account.AccountNumber = req.AccountNumber
	account.FirstName = req.FirstName
	account.LastName = req.LastName

	// call the create account with that struct
	acc, err := s.store.CreateAccount(account)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, acc)
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

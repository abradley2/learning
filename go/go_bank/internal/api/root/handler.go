package root

import (
	"database/sql"

	"github.com/wley3337/learning/tree/main/go/go_bank/internal/handler"
	t "github.com/wley3337/learning/tree/main/go/go_bank/types"

	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	DB *sql.DB
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	s := defaultStore{h.DB}

	switch method := r.Method; method {
	case "GET":
		err = handleGetAccounts(s, w, r)
	case "POST":
		err = handleCreateAccount(s, w, r)
	default:
		err = fmt.Errorf("Method not allowed %s", r.Method)
	}

	if err != nil {
		handler.WriteJSONError(w, http.StatusInternalServerError, err)
	}
}

func handleGetAccounts(s Store, w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.GetAccounts()
	if err != nil {
		log.Println("Error getting accounts")
		return err
	}
	return handler.WriteJSON(w, http.StatusOK, accounts)
}

type CreateAccountRequestBody struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber uuid.UUID `json:"accountNumber"`
	Balance       int64     `json:"balance"`
}

func handleCreateAccount(s Store, w http.ResponseWriter, r *http.Request) error {
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
	account := t.CreateAccountAccount{}
	account.AccountNumber = req.AccountNumber
	account.FirstName = req.FirstName
	account.LastName = req.LastName

	// call the create account with that struct
	acc, err := s.CreateAccount(account)

	if err != nil {
		return err
	}

	return handler.WriteJSON(w, http.StatusCreated, acc)
}

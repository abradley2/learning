package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// write our response to json
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// once you run `WriteHeader`, you can no longer update it
	w.Header().Set("Content-Type", "application/json")
	// WriteHeader locks header
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// wrapper to allow for handling of errors acts as a decorator
type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
			// handle error here
		}
	}
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/accounts/{id}", makeHTTPHandleFunc(s.handleGetAccount))
	log.Println("JSON API server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch method := r.Method; method {
	case "GET":
		return s.handleGetAccounts(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("Method not allowed %s", r.Method)

	}
}

// list
func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	// account := NewAccount("Money", "Bags")
	// account1 := NewAccount("Money", "Penny")
	// account2 := NewAccount("James", "Bond")
	// accounts := []*Account{account, account1, account2}
	accounts := []*Account{}
	return WriteJSON(w, http.StatusOK, accounts)
}

// detail
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	// doesn't error if no id provided, just empty string
	ID := mux.Vars(r)["id"]
	println("ID %s", ID)
	account := NewAccount("Money", "Bags")
	return WriteJSON(w, http.StatusOK, account)
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
	account := new(CreateAccountAccount)
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

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

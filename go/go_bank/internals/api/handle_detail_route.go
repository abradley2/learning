package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleAccountDetails(w http.ResponseWriter, r *http.Request) error {
	switch method := r.Method; method {
	case "GET":
		return s.handleGetAccountByID(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("Method not allowed %s", r.Method)

	}
}

// detail
func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	// doesn't error if no id provided, just empty string
	ID := mux.Vars(r)["id"]
	log.Println("getting account by id:",
		ID)

	id, err := strconv.Atoi(ID)
	if err != nil {
		log.Println("Error converting account ID to string")
		return fmt.Errorf("Invalid ID given %s", ID)
	}
	// call the create account with that struct
	acc, err := s.store.GetAccountByID(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, acc)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	ID := mux.Vars(r)["id"]
	log.Println("deleting account by id:",
		ID)

	id, conversion_err := strconv.Atoi(ID)
	if conversion_err != nil {
		log.Println("Error converting account ID to string")
		return fmt.Errorf("Invalid ID given %s", ID)
	}

	err := s.store.DeleteAccount(id)

	if err != nil {
		return err
	}

	response := "Deleted account with id: " + ID
	return WriteJSON(w, http.StatusOK, response)
}

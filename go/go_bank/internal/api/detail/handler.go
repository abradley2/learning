package detail

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wley3337/learning/tree/main/go/go_bank/internal/handler"
)

type Handler struct {
	DB *sql.DB
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	s := defaultStore{h.DB}

	switch method := r.Method; method {
	case "GET":
		err = handleGetAccountByID(s, w, r)
	case "DELETE":
		err = handleDeleteAccount(s, w, r)
	default:
		err = fmt.Errorf("Method not allowed %s", r.Method)
	}

	if err != nil {
		handler.WriteJSONError(w, http.StatusInternalServerError, err)
	}
}

// detail
func handleGetAccountByID(s Store, w http.ResponseWriter, r *http.Request) error {
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
	acc, err := s.GetAccountByID(id)

	if err != nil {
		return err
	}

	return handler.WriteJSON(w, http.StatusOK, acc)
}

func handleDeleteAccount(s Store, w http.ResponseWriter, r *http.Request) error {
	ID := mux.Vars(r)["id"]
	log.Println("deleting account by id:",
		ID)

	id, conversion_err := strconv.Atoi(ID)
	if conversion_err != nil {
		log.Println("Error converting account ID to string")
		return fmt.Errorf("Invalid ID given %s", ID)
	}

	err := s.DeleteAccount(id)

	if err != nil {
		return err
	}

	response := "Deleted account with id: " + ID
	return handler.WriteJSON(w, http.StatusOK, response)
}

package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wley3337/learning/tree/main/go/go_bank/internal/api/detail"
	"github.com/wley3337/learning/tree/main/go/go_bank/internal/api/root"
)

// wrapper to allow for handling of errors acts as a decorator
type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func RunServer(addr string, db *sql.DB) error {
	router := mux.NewRouter()
	router.Handle("/accounts/{id}", detail.Handler{DB: db})
	router.Handle("/accounts", root.Handler{DB: db})
	log.Println("JSON API server running on port:", addr)
	return http.ListenAndServe(addr, router)
}

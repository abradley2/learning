package api

import (
	"encoding/json"
	st "github.com/wley3337/learning/tree/main/go/go_bank/internals/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// wrapper to allow for handling of errors acts as a decorator
type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}
type APIServer struct {
	listenAddr string
	store      st.Storage
}

// write our response to json
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// once you run `WriteHeader`, you can no longer update it
	w.Header().Set("Content-Type", "application/json")
	// WriteHeader locks header
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
			// handle error here
		}
	}
}

func NewAPIServer(listenAddr string, store st.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}", makeHTTPHandleFunc(s.handleAccountDetails))
	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.handleAccount))
	log.Println("JSON API server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

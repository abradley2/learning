package handler

import (
	"encoding/json"
	"net/http"
)

func WriteJSONError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

// write our response to json
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// once you run `WriteHeader`, you can no longer update it
	w.Header().Set("Content-Type", "application/json")
	// WriteHeader locks header
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

package main

import (
	"github.com/wley3337/learning/tree/main/go/go_bank/internals/api"
	st "github.com/wley3337/learning/tree/main/go/go_bank/internals/store"
	"log"
)

func main() {
	store, err := st.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}

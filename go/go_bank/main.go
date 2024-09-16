package main

import (
	"log"

	"github.com/wley3337/learning/tree/main/go/go_bank/internal/api"
	"github.com/wley3337/learning/tree/main/go/go_bank/internal/database"
)

func main() {
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Setup(db); err != nil {
		log.Fatal(err)
	}

	err = api.RunServer(":3000", db)
	if err != nil {
		_ = db.Close()
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/nirwan77/goLang/cmd/api"
	"github.com/nirwan77/goLang/cmd/db"
)

func main() {
	dbConnection, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	var conn = api.NewAPIServer(":8000", dbConnection)
	if err := conn.Run(); err != nil {
		log.Fatal("Error connecting")
	}
}

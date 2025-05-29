package main

import (
	"log"

	"github.com/nirwan77/goLang/cmd/api"
)

func main() {
	var conn = api.NewAPIServer(":8000", nil)
	if err := conn.Run(); err != nil {
		log.Fatal("Error connecting")
	}
}

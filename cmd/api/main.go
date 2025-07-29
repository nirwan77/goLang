package main

import (
	"database/sql"
	"log"
	"res-api-gin/internal/database"
	"res-api-gin/internal/env"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type application struct {
	port      int
	jswSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:admin@localhost:5432/GoDatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jswSecret: env.GetEnvString("JWT_SECRET", "some-secret"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

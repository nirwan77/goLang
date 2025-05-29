package api

import (
	"database/sql"
	"fmt"
	"net/http"

	handlers "github.com/nirwan77/goLang/cmd/handler"
)

type APIServer struct {
	Address string
	db      *sql.DB
}

func NewAPIServer(Address string, db *sql.DB) *APIServer {
	return &APIServer{
		Address,
		db,
	}
}

func (s *APIServer) routes() {
	http.HandleFunc("/pokemon", handlers.HandlePostPokemon())
}

func (s *APIServer) Run() error {
	s.routes()
	fmt.Println("Server is running in: ", s.Address)
	return http.ListenAndServe(s.Address, nil)
}

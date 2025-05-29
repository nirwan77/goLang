package api

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	handlers "github.com/nirwan77/goLang/cmd/handler"
)

type APIServer struct {
	Address string
	db      *pgx.Conn
}

func NewAPIServer(Address string, db *pgx.Conn) *APIServer {
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

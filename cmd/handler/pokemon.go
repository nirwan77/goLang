package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func HandleGetPokemon(s *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rows, err := s.Query(r.Context(), "SELECT id, name, type FROM pokemon")
			if err != nil {
				return
			}
			defer rows.Close()

			pokemons, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Pokemon, error) {
				var p Pokemon
				err := row.Scan(&p.ID, &p.Name, &p.Type)
				return p, err
			})

			if err != nil {
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pokemons)
		}

		if r.Method == http.MethodPost {
			var updateData Pokemon

			if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
				fmt.Print(err)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			query := `INSERT INTO pokemon (name, type) VALUES ($1, $2)`

			_, err := s.Exec(r.Context(), query, updateData.Name, updateData.Type)
			if err != nil {
				fmt.Print(err)
				return
			}

			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Pokemon inserted successfully",
				"data":    updateData,
			})
		}
	}
}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
			rows, err := s.Query(r.Context(), "SELECT id, name, type FROM pokemon order by id")
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

		if r.Method == http.MethodPatch {
			var patchData Pokemon

			if err := json.NewDecoder(r.Body).Decode(&patchData); err != nil {
				fmt.Print(err)
				return
			}

			query := "Update Pokemon set name=$1, type=$2 where id = $3"

			res, err := s.Exec(r.Context(), query, patchData.Name, patchData.Type, patchData.ID)
			if err != nil {
				fmt.Println(err)
			}

			if res.RowsAffected() == 0 {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Pokemon with given ID not found",
				})
				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": fmt.Sprintf("successfully updated id %d", patchData.ID),
				"data": map[string]interface{}{
					"name": patchData.Name,
					"type": patchData.Type,
				},
			})

		}

		if r.Method == http.MethodPost {
			var updateData Pokemon

			if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
				fmt.Print(err)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			tx, err := s.Begin(r.Context())
			if err != nil {
				fmt.Println(err)
				return
			}

			defer tx.Rollback(r.Context())

			query := `INSERT INTO pokemon (name, type) VALUES ($1, $2)`

			_, err = tx.Exec(r.Context(), query, updateData.Name, updateData.Type)
			if err != nil {
				fmt.Print(err)
				return
			}

			tx.Commit(r.Context())

			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Pokemon inserted successfully",
				"data":    map[string]interface{}{"name": updateData.Name, "type": updateData.Type},
			})
		}
	}
}

func HandleDeletePokemon(s *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Method == http.MethodDelete {
			parts := strings.Split(r.URL.Path, "/")
			if len(parts) != 3 || parts[2] == "" {
				http.Error(w, "Invalid or missing ID in URL", http.StatusBadRequest)
				return
			}

			_, err := s.Exec(r.Context(), "Delete from pokemon where id = $1", parts[2])
			if err != nil {
				fmt.Print(err)
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": fmt.Sprintf("Id no %s was deleted successfully", parts[2]),
			})
		}
	}
}

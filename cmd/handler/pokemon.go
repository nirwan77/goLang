package handler

import (
	"encoding/json"
	"net/http"
)

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func HandleGetPokemon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pokemons := []Pokemon{
				{ID: 1, Name: "Bulbasaur", Type: "Grass"},
				{ID: 2, Name: "Charmander", Type: "Fire"},
				{ID: 3, Name: "Squirtle", Type: "Water"},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pokemons)
		}

		if r.Method == http.MethodPost {
			var updateData Pokemon

			if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
				http.Error(w, "Invalid JSON body", http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Pokemon updated successfully",
				"data":    updateData,
			})
		}
	}
}

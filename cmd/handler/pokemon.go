package handler

import (
	"fmt"
	"net/http"
)

func HandlePostPokemon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a = map[string]string{"1": "Hero"}
		a["asd"] = "asjdlk"

		if r.Method == http.MethodGet {
			fmt.Fprintln(w, a)
		} else {
			fmt.Fprintln(w, "Pokémon created!")
		}
	}
}

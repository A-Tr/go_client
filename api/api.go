package api

import (
	"net/http"
	"pokego/client"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	pokemonUrl = "https://pokeapi.co/api/v2/pokemon/"
)

func HandlePokemon(w http.ResponseWriter, r *http.Request) {
	log.Print("Received request")

	vars := mux.Vars(r)
	pokeId := vars["id"]
	log.Print("Pokemon Id to look for: ", pokeId)

	client := client.InitClient()
	body, code, err := client.GetUrl(pokemonUrl, pokeId)
	if err != nil {
		w.WriteHeader(code)
		return
	}

	w.WriteHeader(code)
	w.Write(body)
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{id}", HandlePokemon)
	return router
}

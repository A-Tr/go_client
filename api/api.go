package api

import (
	"pokego/client"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
)


func HandlePokemon(w http.ResponseWriter, r *http.Request) {
	log.Print(w, "On my Way!")

	vars := mux.Vars(r)
    pokeId := vars["id"]
    log.Print(w, "Pokemon Id: ", pokeId)
	
	client := client.InitClient()
	body, err := client.GetSinglePokemon(pokeId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func InitRouter () *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/{id}", HandlePokemon)
	return router
}
package main

import (
	"net/http"
	"pokego/api"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Print("Initializing server")
}

func main() {

	router := api.InitRouter()
	log.Fatal(http.ListenAndServe(":3001", router))
}

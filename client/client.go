package client

import (
	"os"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"fmt"
	"net/http"
)

type PokeClient interface {
	GetSinglePokemon(string) ([]byte, error)
} 

type RealClient struct {}


func (c RealClient) GetSinglePokemon(id string) ([]byte, error) {
	url := fmt.Sprint(pokeURL, "/pokemon/", id)

	log.Print("Doing request to: ", url)
	resp, err := http.Get(pokeURL)
	if err != nil {
		log.Error("Error doing the request: ", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error doing the request: ", err)
		return nil, err
	}
	
	return body, nil
}


type TestClient struct {}

func (t TestClient) GetSinglePokemon(id string) ([]byte, error) {
	log.Print("I'm the fake client bruh")
	return []byte("1"), nil
}

func InitClient() PokeClient {
	if (os.Getenv("ENV") == "TEST") {
		return new(TestClient)
	}
	return new(RealClient)
}
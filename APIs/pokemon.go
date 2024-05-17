package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type responseHttp struct {
	Nome    string    `json:"name"`
	Pokemon []pokemon `json:"pokemon_entries"`
}
type pokemon struct {
	Numero  int            `json:"entry_number"`
	Especie pokemonSpecies `json:"pokemon_species"`
}

type pokemonSpecies struct {
	Nome string `json:"name"`
}

//http://pokeapi.co/api/v2/pokedex/kanto/
func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)

	}

	//fmt.Println(string(responseData))

	var responseObject responseHttp

	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Nome)
	fmt.Println(responseObject.Pokemon)
	fmt.Println(responseObject.Pokemon[1].Especie)

}

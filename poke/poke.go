// usaremos a pokeapi (lá tem tudo sobre pokemon)
// para consultar: http:// pokeapi.co/api/v2/pokemon/kanto
// vamos usar o verbo GET

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mhishmeh/pokedexcli/internal/pokecache"
)

type CatchPokemon struct {
	Name    string `json:"name"`
	BaseEXP int    `json:"base_experience"`
}
type PokeEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	}
}
type PokeArea struct {
	Location          string          `json:"name"`
	PokemonEncounters []PokeEncounter `json:"pokemon_encounters"`
}
type PokeApiResponse struct {
	Results  []PokeArea `json:"results"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
}

var URL = "https://pokeapi.co/api/v2/location-area"
var PreviousURL string
var cache = pokecache.NewCache(20 * time.Second)
var PokemonURL = "https://pokeapi.co/api/v2/pokemon/"

var commands = map[string]func(args ...string){
	"help": func(args ...string) {
		fmt.Println("Hello! The help function exists to show you availabe commands! They are currently \nhelp: shows available commands\n exit:quits the REPL.")
	},

	"exit": func(args ...string) {
		fmt.Println("Exiting REPL. Goodbye!")
	},

	"map": func(args ...string) {
		if data, ok := cache.Get(URL); ok {
			// Use cached data
			var locations PokeApiResponse
			if err := json.Unmarshal(data, &locations); err != nil {
				log.Fatalf("couldn't parse cached data: %v", err)
			}
			for _, location := range locations.Results {
				cleanedLocation := strings.ReplaceAll(location.Location, "{", "")
				cleanedLocation = strings.ReplaceAll(cleanedLocation, "}", "")
				fmt.Println(cleanedLocation) // Print without braces
			}
			URL = locations.Next
			PreviousURL = locations.Previous
		}
		req, err := http.Get(URL)
		if err != nil {
			log.Fatalf("couldnt do %v", err)
		}
		defer req.Body.Close()
		var locations PokeApiResponse
		if err := json.NewDecoder(req.Body).Decode(&locations); err != nil {
			log.Fatalf("couldnt do %v", err)
		}
		for _, location := range locations.Results {
			cleanedLocation := strings.ReplaceAll(location.Location, "{", "")
			cleanedLocation = strings.ReplaceAll(cleanedLocation, "}", "")
			fmt.Println(cleanedLocation) // Print without braces
		}
		cachedData, err := json.Marshal(locations)
		if err != nil {
			log.Fatalf("couldn't marshal data to cache: %v", err)
		}
		cache.Add(URL, cachedData)
		URL = locations.Next
		PreviousURL = locations.Previous

	},
	"mapb": func(args ...string) {
		if PreviousURL == "" {
			fmt.Println("Sorry, this is the first page.")
			return
		}

		req, err := http.Get(PreviousURL)
		if err != nil {
			log.Fatalf("couldnt do %v", err)
		}
		defer req.Body.Close()
		var locations PokeApiResponse
		if err := json.NewDecoder(req.Body).Decode(&locations); err != nil {
			log.Fatalf("couldnt do %v", err)
		}
		for _, location := range locations.Results {
			cleanedLocation := strings.ReplaceAll(location.Location, "{", "")
			cleanedLocation = strings.ReplaceAll(cleanedLocation, "}", "")
			fmt.Println(cleanedLocation) // Print without braces
		}
		URL = locations.Next
		PreviousURL = locations.Previous
	},

	"explore": func(name ...string) {
		// Make the URL dynamic based on the location name
		url := URL + "/" + name[0]
		fmt.Printf("Fetching data from: %s\n", url) // Debugging: Print the URL being requested

		// Make the GET request
		req, err := http.Get(url)
		if err != nil {
			log.Fatalf("Couldn't fetch data for %s: %v", name, err)
		}
		defer req.Body.Close()

		// Debugging: Print the raw response body

		// You can stop here for debugging purposes and examine the structure of `rawResponse`
		// Decode the response into the expected structure if the API response is correct
		var pokemon PokeArea
		if err := json.NewDecoder(req.Body).Decode(&pokemon); err != nil {
			log.Fatalf("Couldn't decode response: %v", err)
		}

		// Display the results (assuming the API response matches the PokeApiResponse structure)
		fmt.Printf("Pokémon found in %s:\n", name[0])
		for _, pokemon := range pokemon.PokemonEncounters {
			fmt.Println(pokemon.Pokemon.Name)
		}
	},
	"catch": func(name ...string) {
		url := PokemonURL + "/" + name[0]
		req, err := http.Get(url)
		if err != nil {
			log.Fatalf("couldnt do %v", err)
		}
		defer req.Body.Close()
		var pokemon CatchPokemon
		if err := json.NewDecoder(req.Body).Decode(&pokemon); err != nil {
			log.Fatalf("couldnt do %v", err)
		}
		fmt.Println("throwing a ball at said pokemon")
		fmt.Println(pokemon.Name, "has been caught!")
		fmt.Printf("%s's exp is %v", pokemon.Name, pokemon.BaseEXP)

	},
}

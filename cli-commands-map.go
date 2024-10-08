package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PokeArea struct {
	Location string `json:"name"`
}
type PokeApiResponse struct {
	Results  []PokeArea `json:"results"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
}

var URL = "https://pokeapi.co/api/v2/location-area"
var PreviousURL string
var commands = map[string]func(){
	"help": func() {
		fmt.Println("Hello! The help function exists to show you availabe commands! They are currently \nhelp: shows available commands and exit.")
	},

	"exit": func() {
		fmt.Println("Exiting REPL. Goodbye!")
	},

	"map": func() {
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
		URL = locations.Next
		PreviousURL = locations.Previous

	},
	"mapb": func() {
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
}

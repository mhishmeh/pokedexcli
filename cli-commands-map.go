package main

import "fmt"

var commands = map[string]func(){
	"help": func() {
		fmt.Println("Hello! The help function exists to show you availabe commands! They are currently \nhelp: shows available commands and exit.")
	},

	"exit": func() {
		fmt.Println("Exiting REPL. Goodbye!")
	},
}

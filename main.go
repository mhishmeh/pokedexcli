package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if command, exists := commands[input]; exists {
			command()

		} else {
			fmt.Println("Sorry this command does not exist. Try these: ")
			for key := range commands {
				fmt.Println(key)
			}
		}
		if input == "exit" {
			break
		}

	}

}

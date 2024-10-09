package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// REPL loop
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Split input into command and arguments
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		cmd := parts[0]
		args := parts[1:]

		// Execute command if it exists

		if command, found := commands[cmd]; found {
			command(args...)
		} else {
			fmt.Println("Sorry this command does not exist. Try these: ")
			for key := range commands {
				fmt.Println(key)
			}
		}
		if cmd == "exit" {
			break
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter info please: ")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		//input := scanner.Text()

	}

}

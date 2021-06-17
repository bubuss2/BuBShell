package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// input reader
	reader := bufio.NewReader(os.Stdin)

	for {
		currentPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		fmt.Print(currentPath)
		fmt.Print(" > ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = executeInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// input reader
	reader := bufio.NewReader(os.Stdin)

	// color variables for colored output
	colorGreen := "\033[32m"
	colorRed := "\033[31m"
	colorDefault := "\033[0m"

	for {
		currentPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		fmt.Print(string(colorGreen), currentPath)
		fmt.Print(string(colorRed), " > ", string(colorDefault))

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		pipeline := strings.Contains(input, "|")

		if pipeline {
			pipelineHandling(input)
		} else {
			if err = executeInput(input); err != nil {
				fmt.Fprintln(os.Stderr, "Command not found!")
			}
		}
	}
}

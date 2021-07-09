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

		if input == "" {
			continue
		}

		multipleCommands := strings.Contains(input, "&&")

		if multipleCommands {
			commands := strings.Split(input, "&&")
			for index := range commands {
				commands[index] = strings.TrimSpace(commands[index])
				pipeline := strings.Contains(commands[index], "|")
				if pipeline {
					executePipelineInput(commands[index])
				} else {
					executeInput(commands[index])
				}
			}
		} else {
			pipeline := strings.Contains(input, "|")
			if pipeline {
				executePipelineInput(input)
			} else {
				executeInput(input)
			}
		}
	}
}

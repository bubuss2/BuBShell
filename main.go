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

		multiple_commands := strings.Contains(input, "&&")

		if multiple_commands {
			multiple_input := strings.Split(input, "&&")
			for index := range multiple_input {
				multiple_input[index] = strings.TrimSpace(multiple_input[index])
				pipeline := strings.Contains(multiple_input[index], "|")
				if pipeline {
					pipelineHandling(multiple_input[index])
				} else {
					executeInput(multiple_input[index])
				}
			}
		} else {
			pipeline := strings.Contains(input, "|")
			if pipeline {
				pipelineHandling(input)
			} else {
				executeInput(input)
			}
		}
	}
}

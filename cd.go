package main

import (
	"log"
	"os"
)

func cd(argument string) error {
	if argument == "$HOME" {
		homePath, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		return os.Chdir(homePath)
	}

	return os.Chdir(argument)
}

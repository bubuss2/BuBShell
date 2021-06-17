package main

import (
	"errors"
	"log"
	"os"
)

func cd(arguments []string) error {
	if len(arguments) < 2 {
		return errors.New("invalid path")
	}

	if arguments[1] == "$HOME" {
		homePath, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		return os.Chdir(homePath)
	}

	return os.Chdir(arguments[1])
}

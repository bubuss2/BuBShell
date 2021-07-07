package commands

import (
	"fmt"
	"log"
	"os"
)

func ChangeDirectory(arguments []string) {
	if len(arguments) < 2 {
		fmt.Fprintln(os.Stderr, "Invalid path!")
		return
	}

	if arguments[1] == "$HOME" {
		homePath, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		if err := os.Chdir(homePath); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		if err := os.Chdir(arguments[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

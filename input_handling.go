package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func executeInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	arguments := strings.Split(input, " ")

	switch arguments[0] {
	case "cd":
		if len(arguments) < 2 {
			return errors.New("invalid path")
		}

		return cd(arguments[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(arguments[0], arguments[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

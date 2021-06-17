package main

import (
	"os"
	"os/exec"
	"strings"
)

func executeInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	arguments := strings.Split(input, " ")

	cmd := exec.Command(arguments[0], arguments[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

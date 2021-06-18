package main

import (
	"os"
	"os/exec"
	"strings"
)

func executeInput(input string) error {
	input = strings.TrimSpace(input)
	arguments := strings.Split(input, " ")

	switch arguments[0] {
	case "top":
		cmd := exec.Command("top", "-b", "-n", "1")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()

	case "cd":
		return cd(arguments)

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(arguments[0], arguments[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

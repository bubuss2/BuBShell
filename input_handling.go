package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"./commands"
)

func executeInput(input string) {
	input = strings.TrimSpace(input)
	arguments := strings.Split(input, " ")

	switch arguments[0] {
	case "top":
		commands.Top(arguments)
		return
	case "cd":
		commands.ChangeDirectory(arguments)
		return
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(arguments[0], arguments[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func pipelineHandling(input string) {
	commands := strings.Split(input, "|")

	for index := range commands {
		commands[index] = strings.TrimSpace(commands[index])
	}

	var arguments [][]string

	for _, command := range commands {
		arguments = append(arguments, strings.Split(command, " "))
	}

	var outputBuffer bytes.Buffer
	for index := range commands[:len(commands)-1] {
		cmd1 := exec.Command(arguments[index][0], arguments[index][1:]...)
		cmd2 := exec.Command(arguments[index+1][0], arguments[index+1][1:]...)
		read, write := io.Pipe()
		cmd1.Stdout = write
		cmd2.Stdin = read

		var tempBuffer bytes.Buffer
		cmd2.Stdout = &tempBuffer

		if err := cmd1.Start(); err != nil {
			if err = cmd1.Wait(); err != nil {
				fmt.Fprintln(os.Stderr, arguments[index][0]+": command not found")
				return
			}
		}

		if err := cmd2.Start(); err != nil {
			if err = cmd2.Wait(); err != nil {
				fmt.Fprintln(os.Stderr, arguments[index+1][0]+": command not found")
				return
			}
		}

		if err := cmd1.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err := write.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err := cmd2.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		outputBuffer = tempBuffer
	}
	io.Copy(os.Stdout, &outputBuffer)
}

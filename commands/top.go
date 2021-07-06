package commands

import (
	"os"
	"os/exec"
)

func Top(arguments []string) error {
	top := []string{"top", "-b", "-n", "1"}
	arguments = append(top, arguments[1:]...)

	cmd := exec.Command(arguments[0], arguments[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

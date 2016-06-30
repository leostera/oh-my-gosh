//+build windows

package omg

import (
	"os/exec"
)

// ClearCmd is the command used to clean the terminal
const ClearCmd = "cls"

func makeCmd(command []string) *exec.Cmd {
	return exec.Command("cmd", append([]string{"/C"}, command...)...)
}

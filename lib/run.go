package omg

import (
	"os"
	"os/exec"
	"syscall"
)

func Reset() {
	Run([]string{ClearCmd})
}

func MakeCmd(command []string) *exec.Cmd {
	cmd := _makeCmd(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func GetExitStatus(err error) int {
	if err == nil {
		return 0
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		if s, ok := exitErr.ProcessState.Sys().(syscall.WaitStatus); ok {
			return s.ExitStatus()
		}
	}
	return -1
}

func CommandExists(command []string) bool {
	_, result := exec.LookPath(command[0])
	return result == nil
}

func Run(command []string) int {
	cmd := MakeCmd(command)
	return GetExitStatus(cmd.Run())
}

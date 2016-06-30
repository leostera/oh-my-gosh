package omg

import (
	"fmt"
	"os"
)

func Die(status int, message string) {
	if len(message) > 0 {
		fmt.Println(message)
	}
	os.Exit(status)
}

func PrintStatus(s int) {
	fmt.Printf("\033[90mexit: %d\033[0m\n\n", s)
}

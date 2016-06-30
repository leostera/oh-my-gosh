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

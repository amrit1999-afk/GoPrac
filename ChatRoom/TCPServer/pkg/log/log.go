package log

import "fmt"

func Logger(origin string) func(string) {
	log := func(message string) {
		messageFromOrigin := "[" + origin + "] " + message
		fmt.Println(messageFromOrigin)
	}

	return log
}

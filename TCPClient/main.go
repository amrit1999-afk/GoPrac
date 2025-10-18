package main

import (
	"fmt"

	"github.com/debdut/TCPClient/pkg/tcp"
)

func main() {
	conn, err := tcp.Connect("localhost:5443")
	if err != nil {
		fmt.Println(err)
		return
	}
	messages := []string{
		"hello",
		"how are you doing?",
		"Great!",
	}

}

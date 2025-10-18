package main

import (
	"fmt"
	"time"

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

	for _, msg := range messages {
		fmt.Println("Message:", msg)
		reply, err := tcp.Message(conn, msg+"\r\n")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Reply:", reply)
		}
		time.Sleep(2 * time.Second)
	}

}

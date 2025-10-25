package main

import (
	"fmt"
	"time"

	"github.com/debdut/TCPClient/pkg/tcp"
)

func main() {
	clientID := tcp.UniqueIDGenerator()
	fmt.Println("Client ID: ", clientID)
	conn, err := tcp.Connect("localhost:5443", clientID)
	if err != nil {
		fmt.Println(err)
		return
	}

	messages := []string{
		"hello Maakichu\n",
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

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/amrit/TCPClient/pkg/tcp"
)

func main() {
	clientID := tcp.UniqueIDGenerator()
	conn, err := tcp.Connect("localhost:5443", clientID)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)

	go tcp.ReceiveMessage(conn, reader)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		fmt.Fprintf(conn, "%s\n", text)
	}
}

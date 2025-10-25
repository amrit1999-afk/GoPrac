package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/amrit/TCPClient/pkg/receive"
	"github.com/amrit/TCPClient/pkg/tcp"
	"github.com/amrit/TCPClient/pkg/uid"
)

func main() {
	clientID := uid.UniqueIDGenerator()
	conn, err := tcp.Connect("localhost:5443", clientID)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)

	go receive.ReceiveMessage(conn, reader)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		fmt.Fprintf(conn, "%s\n", text)
	}
}

package receive

import (
	"bufio"
	"fmt"
	"net"
)

func ReceiveMessage(conn net.Conn, reader *bufio.Reader) {
	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Disconnected from server")
			return
		}

		fmt.Println(msg)
	}
}

package tcp

import (
	"fmt"
	"net"

	"github.com/amrit/TCPClient/pkg/log"
)

var clientLog func(string)

func Connect(address string, clientID string) (net.Conn, error) {
	clientLog = log.Logger("TCP Client")
	conn, err := net.Dial("tcp", address)

	if err != nil {
		clientLog("Failed to connect with " + address)
	} else {
		fmt.Fprintf(conn, "%s\n", clientID)
		clientLog("Connected with " + address)
	}

	return conn, err
}

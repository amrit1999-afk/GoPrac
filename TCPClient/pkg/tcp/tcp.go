package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	"github.com/debdut/TCPClient/pkg/log"
)

var clientLog func(string)

func HTTPStatusCheck(addr string) (string, error) {
	conn, err := Connect(addr)
	if err != nil {
		return "", err
	}

	status, err := Message(conn, "GET / HTTP/1.0/r/n/r/n")
	if err != nil {
		return "", err
	}

	return status, err
}

func Connect(addr string) (net.Conn, error) {
	clientLog = log.Logger("TCP Client")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		clientLog("Failed to connect with " + addr)
	} else {
		clientLog("Connected with " + addr)
	}

	return conn, err
}

func Message(conn net.Conn, message string) (string, error) {
	n, err := fmt.Fprint(conn, message)
	if err != nil {
		clientLog("Failed to write")
		return "", err
	}
	clientLog("Wrote " + strconv.Itoa(n) + " bytes")

	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		clientLog("Couldn't read")
		return "", err
	}
	clientLog("Message read")

	return reply, nil
}

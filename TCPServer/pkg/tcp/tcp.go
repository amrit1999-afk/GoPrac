package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/debdut/TCPServer/pkg/log"
)

var serverLog func(string)

func StartServer(addr string) error {
	serverLog = log.Logger("TCP Server")

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		serverLog("Failed to start")
	}
	serverLog("Started at: " + addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			serverLog("Couldn't accept connection")
		} else {
			serverLog("Accepted Connection")
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	serverLog("Conn handler")
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			serverLog("Disconnected")
			break
		}

		serverLog("Message: " + strings.TrimSpace(msg))
		fmt.Fprintf(conn, msg+"\r\n")
	}
}

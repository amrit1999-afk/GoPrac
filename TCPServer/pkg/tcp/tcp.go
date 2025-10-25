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
			reader := bufio.NewReader(conn)
			cliendID, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Failed to get Client ID")
				conn.Close()
				continue
			}
			serverLog("Successfully Connected with Client: " + cliendID)
			go handleConn(conn, reader)
			// handleConn(conn, reader)
		}
	}
}

func handleConn(conn net.Conn, reader *bufio.Reader) {
	defer conn.Close()
	serverLog("Conn handler")

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

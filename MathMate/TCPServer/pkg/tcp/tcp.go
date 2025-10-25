package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/amrit/TCPServer/pkg/log"
	"github.com/amrit/TCPServer/pkg/parser"
)

type Message struct {
	SenderID   string
	MsgContent string
}

var (
	ServerLog        func(string)
	clientMap        = make(map[string]net.Conn)
	clientMapMutex   sync.Mutex
	broadcastChannel = make(chan Message)
)

func StartServer(address string) {
	ServerLog = log.Logger("TCP Server")

	ln, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("Failed to start server ", err)
		return
	}
	ServerLog("Server started successfully")
	go broadcast()
	acceptClients(ln)
}

func acceptClients(ln net.Listener) {
	for {
		conn, err := ln.Accept()

		if err != nil {
			ServerLog("Failed to establish Connection")
		} else {
			reader := bufio.NewReader(conn)

			clientID, err := reader.ReadString('\n')

			if err != nil {
				ServerLog("Failed to get client ID")
				conn.Close()
				continue
			}

			ServerLog("Successfully connected with client: " + clientID)
			addClient(clientID, conn, reader)
		}
	}
}

func addClient(clientID string, conn net.Conn, reader *bufio.Reader) {
	clientID = strings.TrimSpace(clientID)
	clientMapMutex.Lock()
	clientMap[clientID] = conn
	clientMapMutex.Unlock()
	go handleClientMsg(clientID, conn, reader)
}

func handleClientMsg(clientID string, conn net.Conn, reader *bufio.Reader) {
	defer conn.Close()
	for {

		msg, err := reader.ReadString('\n')

		if err != nil {
			ServerLog("Connection lost with Client: " + clientID)
			clientMapMutex.Lock()
			delete(clientMap, clientID)
			clientMapMutex.Unlock()
			return
		}

		if msg == "" {
			continue
		}

		msg = strings.TrimSpace(msg)
		broadcastChannel <- Message{SenderID: clientID, MsgContent: msg}
	}
}

func broadcast() {
	for {
		msg := <-broadcastChannel
		expression := msg.MsgContent
		expVal, err := parser.EvaluateExpression(expression)

		if err != nil {
			ServerLog("Error evaluating expression " + err.Error())
			continue
		}
		clientMapMutex.Lock()

		conn, found := clientMap[msg.SenderID]

		clientMapMutex.Unlock()

		if found {
			fmt.Fprintf(conn, "[Server]: Value of [%s] = %d\r\n", msg.MsgContent, expVal)
		} else {
			ServerLog("Client not found: " + msg.SenderID)
		}

	}

}

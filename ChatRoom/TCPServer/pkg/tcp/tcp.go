package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/amrit/TCPServer/pkg/log"
)

type Message struct {
	SenderID   string
	MsgContent string
}

var ServerLog func(string)
var clientMap map[string]net.Conn = make(map[string]net.Conn)
var clientMapMutex sync.Mutex

var broadcastingChannel = make(chan Message)

func StartServer(address string) {
	ServerLog = log.Logger("TCP Server")

	ln, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("Failed to start server ", err)
		return
	}
	ServerLog("Server started successfully")
	go broadcaster()
	acceptClients(ln)
}

func acceptClients(ln net.Listener) {
	for {
		conn, err := ln.Accept()

		if err != nil {
			ServerLog("Failed to connect with client")
		} else {
			reader := bufio.NewReader(conn)
			cliendID, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Failed to get Client ID")
				conn.Close()
				continue
			}
			ServerLog("Successfully Connected with Client: " + cliendID)
			addClient(cliendID, conn, reader)
		}
	}
}

func addClient(cliendID string, conn net.Conn, reader *bufio.Reader) {
	cliendID = strings.TrimSpace(cliendID)
	clientMapMutex.Lock()
	clientMap[cliendID] = conn
	clientMapMutex.Unlock()
	go handleClientMsg(cliendID, conn, reader)
}

func handleClientMsg(clientID string, conn net.Conn, reader *bufio.Reader) {
	defer conn.Close()
	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			ServerLog("Connection lost with Client")
			clientMapMutex.Lock()
			delete(clientMap, clientID)
			clientMapMutex.Unlock()
			return
		}

		if msg == "" {
			continue
		}

		msg = strings.TrimSpace(msg)
		broadcastingChannel <- Message{SenderID: clientID, MsgContent: msg}
	}
}

func broadcaster() {
	for {
		msg := <-broadcastingChannel
		clientMapMutex.Lock()

		for id, conn := range clientMap {

			if id == msg.SenderID {
				continue
			}

			fmt.Fprintf(conn, "%s %s\r\n", msg.SenderID, msg.MsgContent)
		}
		clientMapMutex.Unlock()
	}
}

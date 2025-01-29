package main

import (
	"fmt"
	"log"
	"os"
	"net"
)

// IT MIGHT BE ADDED SOMEWHERE AS ENVIRONMENTAL VARIABLE
const SOCKET_PATH = "/tmp/fiber_orcas.sock"

func connectToSocket() {
	if err := os.Remove(SOCKET_PATH); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	listener, err := net.Listen("unix", SOCKET_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// TO BE DELETED, USED FOR DEVELOPMENT PURPOSES, DEBUG MODE CAN BE ADDED LATER
	fmt.Println("Listening on ", SOCKET_PATH)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				log.Println("Reading error bruv", err)
			}
			return
		}

		data := buffer[:n]
		parsedData, err = parsePacket(data)
		if err != nil {
			log.Println("Parsing error bruv", err)
		}
		outputData(parsedData)
	}
}
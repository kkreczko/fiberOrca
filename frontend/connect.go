package main

import (
	"bufio"
	"github.com/kruczys/fiberOrca/models"
	"log"
	"net"
	"os"
)

const SOCKET_PATH = "/tmp/fiber_orcas.sock"

func connectToSocket(s *models.Session, ch chan string) {
	if err := os.Remove(SOCKET_PATH); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	listener, err := net.Listen("unix", SOCKET_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConnection(conn, s, ch)
	}
}

func handleConnection(conn net.Conn, s *models.Session, ch chan string) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		select {
		case <-ch:
			os.Exit(0)
		default:
			packet, err := parsePacket(scanner.Bytes())
			if err != nil {
				log.Println("Parsing error: ", err)
				continue
			}
			if packet != nil {
				s.Update(packet)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Scanner error:", err)
	}
}

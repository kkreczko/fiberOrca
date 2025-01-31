package main

import (
	"fmt"
	"github.com/kruczys/fiberOrca/models"
)

// TEMPORARY OUTPUT FUNCTION IT WILL BE CHANGED FOR BUBBLETEA INTEGRATION PURPOSES
// FOR NOW IT WORKS
func outputData(packet *models.Packet) {
	if packet == nil {
		return
	}

	fmt.Printf("Protocol: %s\n", packet.Protocol)
	fmt.Printf("Source: %s:%d\n", packet.SourceIP, packet.SourcePort)
	fmt.Printf("Destination: %s:%d\n", packet.DestinationIP, packet.DestinationPort)
	fmt.Printf("Time: %v\n", packet.Datetime())
	fmt.Printf("TTL: %d\n", packet.TTL)
}

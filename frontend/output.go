package main

import (
	"fmt"
)

// TEMPORARY OUTPUT FUNCTION IT WILL BE CHANGED FOR BUBBLETEA INTEGRATION PURPOSES
// FOR NOW IT WORKS
func outputData(packet *Packet) {
	if packet == nil {
		return
	}

	fmt.Printf("Protocol: %s\n", packet.Protocol)
	fmt.Printf("Source: %s:%d\n", packet.SourceIP, packet.SourcePort)
	fmt.Printf("Destination: %s:%d\n", packet.DestIP, packet.DestPort)
	fmt.Printf("Time: %v\n", packet.Timestamp)
	fmt.Printf("TTL: %d\n", packet.TTL)
}

package models

import "time"

// DataLink is a struct that represents the data link layer of a packet
type DataLink struct {
	SourceMAC      string
	DestinationMAC string
	Protocol       string
}

// Network is a struct that represents the network layer of a packet
type Network struct {
	SourceIP      string
	DestinationIP string
	Protocol      string
}

// Transport is a struct that represents the transport layer of a packet
type Transport struct {
	SourcePort      string
	DestinationPort string
	Protocol        string
}

// Packet is a struct that represents a packet
type Packet struct {
	ID        int
	DataLink  DataLink
	Network   Network
	Transport Transport
	TTL       int
	datetime  time.Time
}

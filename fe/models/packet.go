package models

import (
	"github.com/google/uuid"
	"time"
)

// DataLink is a struct that represents the data link layer of a packet
// TODO: Implement the DataLink struct
//type DataLink struct {
//	SourceMAC      string
//	DestinationMAC string
//	Protocol       string
//}

// Network is a struct that represents the network layer of a packet
type Network struct {
	SourceIP      string
	DestinationIP string
}

// Transport is a struct that represents the transport layer of a packet
type Transport struct {
	SourcePort      string
	DestinationPort string
	Protocol        string
}

// Packet is a struct that represents a packet
// It implements the Item interface
type Packet struct {
	ID        uuid.UUID
	Network   Network
	Transport Transport
	TTL       int
	datetime  time.Time
}

func (p Packet) SourceIP() string {
	return p.Network.SourceIP
}

func (p Packet) DestinationIP() string {
	return p.Network.DestinationIP
}

func (p Packet) Protocol() string {
	return p.Transport.Protocol
}

func (p Packet) SourcePort() string {
	return p.Transport.SourcePort
}

func (p Packet) DestinationPort() string {
	return p.Transport.DestinationPort
}

func (p Packet) TTLValue() int {
	return p.TTL
}

func (p Packet) Datetime() time.Time {
	return p.datetime
}

func NewNetwork(sourceIP, destinationIP string) Network {
	return Network{
		SourceIP:      sourceIP,
		DestinationIP: destinationIP,
	}
}

func NewTransport(sourcePort, destinationPort, protocol string) Transport {
	return Transport{
		SourcePort:      sourcePort,
		DestinationPort: destinationPort,
		Protocol:        protocol,
	}
}

func NewPacket(network Network, transport Transport, ttl int, datetime time.Time) Packet {
	return Packet{
		ID:        uuid.New(),
		Network:   network,
		Transport: transport,
		TTL:       ttl,
		datetime:  datetime,
	}
}

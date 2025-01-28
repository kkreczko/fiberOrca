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
// It implements the Item interface
type Packet struct {
	ID        int
	DataLink  DataLink
	Network   Network
	Transport Transport
	TTL       int
	datetime  time.Time
}

func (p Packet) FilterValue() string {
	return p.Network.SourceIP
}

func (p Packet) SourceMac() string {
	return p.DataLink.SourceMAC
}

func (p Packet) DestinationMac() string {
	return p.DataLink.DestinationMAC
}

func (p Packet) DataLinkProtocol() string {
	return p.DataLink.Protocol
}

func (p Packet) SourceIP() string {
	return p.Network.SourceIP
}

func (p Packet) DestinationIP() string {
	return p.Network.DestinationIP
}

func (p Packet) NetworkProtocol() string {
	return p.Network.Protocol
}

func (p Packet) SourcePort() string {
	return p.Transport.SourcePort
}

func (p Packet) DestinationPort() string {
	return p.Transport.DestinationPort
}

func (p Packet) TransportProtocol() string {
	return p.Transport.Protocol
}

func (p Packet) TTLValue() int {
	return p.TTL
}

func (p Packet) Datetime() time.Time {
	return p.datetime
}

func NewPacket(id int, dataLink DataLink, network Network, transport Transport, ttl int, datetime time.Time) Packet {
	return Packet{
		ID:        id,
		DataLink:  dataLink,
		Network:   network,
		Transport: transport,
		TTL:       ttl,
		datetime:  datetime,
	}
}

package main

import (
	"errors"
	"github.com/kruczys/fiberOrca/models"
	"strconv"
	"strings"
	"time"
)

type Packet struct {
	Protocol   string
	SourcePort int
	SourceIP   string
	DestIP     string
	DestPort   int
	Timestamp  time.Time
	TTL        int
}

// THIS FUNCTION PARSES INCOMING PACKETS INTO COOL AND NICE! PACKET TYPE STRUCTURES
// IT JUST WORKS?
func parsePacket(packetData []byte) (*models.Packet, error) {
	data := string(packetData)
	data = strings.TrimRight(data, "\x00")
	data = strings.TrimSpace(data)

	parts := strings.Split(data, ";")

	protocol := parts[0]
	sourceIP := parts[2]
	destIP := parts[3]

	srcPort, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	dstPort, err := strconv.Atoi(parts[4])
	if err != nil {
		return nil, err
	}

	timestampParts := strings.Split(parts[5], ".")
	if len(timestampParts) != 2 {
		return nil, errors.New("invalid timestamp format")
	}

	seconds, err := strconv.ParseInt(timestampParts[0], 10, 64)
	if err != nil {
		return nil, err
	}

	microseconds, err := strconv.ParseInt(timestampParts[1], 10, 64)
	if err != nil {
		return nil, err
	}
	nanoseconds := microseconds * 1000

	ttl, err := strconv.Atoi(parts[6])
	if err != nil {
		return nil, err
	}

	networkProtocol := models.NewNetwork(sourceIP, destIP)
	transportProtocol := models.NewTransport(strconv.Itoa(srcPort), strconv.Itoa(dstPort), protocol)

	packet := models.NewPacket(networkProtocol, transportProtocol, ttl, time.Unix(seconds, nanoseconds))

	return &packet, nil
}

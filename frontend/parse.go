package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Packet struct {
	Protocol 	string
	SourcePort 	int
	SourceIP 	string
	DestIP 		string
	DestPort 	int
	Timestamp 	time.Time
	TTL 		int
}

func parsePacket(packet []byte) (*Packet, error) {
	data = strings.TrimRight(data, "\x00")
	data = strings.TrimSpace(data)

	parts := strings.Split(data, ";")
	if (len(parts) != 7) {
		return nil, errors.New("invalid packet")
	}

	srcPort, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	dstPort, err := strconv.Atoi(parts[4])
	if err != nil {
		return nil, err
	}

	timestamp, err := strconv.ParseInt(parts[5], 10, 64)
	if err != nil {
		return nil, err
	}

	ttl, err := strconv.Atoi(parts[6])
	if err != nil {
		return nil, err
	}

	packet := &Packet {
		Protocol: 	parts[0],
		SourcePort: srcPort,
		SourceIP: 	parts[2],
		DestIP: 	parts[3],
		DestPort: 	dstPort,
		Timestamp: 	time.Unix(timestamp, 0),
		TTL: 		ttl,
	}

	return packet, nil
}
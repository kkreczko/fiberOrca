package main

import (
    "errors"
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

func parsePacket(packetData []byte) (*Packet, error) {
    data := string(packetData)
	data = strings.TrimRight(data, "\x00")
	data = strings.TrimSpace(data)

	parts := strings.Split(data, ";")

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

	packet := &Packet {
		Protocol: 	parts[0],
		SourcePort: srcPort,
		SourceIP: 	parts[2],
		DestIP: 	parts[3],
		DestPort: 	dstPort,
		Timestamp:  time.Unix(seconds, nanoseconds),
		TTL: 		ttl,
	}

	return packet, nil
}
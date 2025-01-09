package models

import (
	"time"
)

// Session is a struct that represents a application session
type Session struct {
	collectedPackets []Packet
	filter           Filter
	startTime        time.Time
	endTime          time.Time
	file             string
	ports            []string
}

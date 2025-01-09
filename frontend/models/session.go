package models

import (
	"time"
)

// Session is a struct that represents a session
type Session struct {
	collectedPackets []Packet
	filter           Filter
	startTime        time.Time
	endTime          time.Time
	file             string
	ports            []string
}

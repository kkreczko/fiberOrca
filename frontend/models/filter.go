package models

import "time"

// Filter is a struct that represents a session filter
type Filter struct {
	MAC                 string
	IP                  string
	senderPort          string
	receiverPort        string
	dataLinkProtocol    string
	networkProtocol     string
	transportProtocol   string
	applicationProtocol string
	startTime           time.Time
	endTime             time.Time
}

// NewFilter creates a new filter
func NewFilter() Filter {
	return Filter{}
}

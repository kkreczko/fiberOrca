package models

import "time"

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

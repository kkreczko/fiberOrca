package models

import (
	"github.com/charmbracelet/bubbles/list"
	"time"
)

func (s *Session) initData(width, height int) {
	packetList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	s.packetPreviews = packetList
	s.packetPreviews.Title = "Packets"
	s.collectedPackets = []Packet{
		{
			ID: 1,
			Network: Network{
				SourceIP:      "1.1.1.1",
				DestinationIP: "2.2.2.2",
				Protocol:      "IPv4",
			},
			Transport: Transport{
				SourcePort:      "80",
				DestinationPort: "80",
				Protocol:        "TCP",
			},
			TTL:      64,
			datetime: time.Now(),
		},
		{
			ID: 2,
			Network: Network{
				SourceIP:      "3.3.3.3",
				DestinationIP: "2.2.2.2",
				Protocol:      "IPv4",
			},
			Transport: Transport{
				SourcePort:      "80",
				DestinationPort: "80",
				Protocol:        "TCP",
			},
			TTL:      64,
			datetime: time.Now(),
		},
		{
			ID: 3,
			Network: Network{
				SourceIP:      "5.5.5.5",
				DestinationIP: "2.2.2.2",
				Protocol:      "IPv4",
			},
			Transport: Transport{
				SourcePort:      "80",
				DestinationPort: "80",
				Protocol:        "TCP",
			},
			TTL:      64,
			datetime: time.Now(),
		},
	}
	for _, packet := range s.collectedPackets {
		s.packetPreviews.InsertItem(-1, list.Item(NewPacketPreview(packet)))
	}

}

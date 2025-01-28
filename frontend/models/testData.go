package models

import (
	"github.com/charmbracelet/bubbles/list"
	"time"
)

func (s *Session) initData(width, height int) {
	packetList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	s.collectedPackets = packetList
	s.collectedPackets.Title = "Packets"
	s.collectedPackets.SetItems([]list.Item{
		&Packet{
			ID: 1,
			DataLink: DataLink{
				SourceMAC:      "00:00:00:00:00:00",
				DestinationMAC: "00:00:00:00:00:00",
				Protocol:       "Ethernet",
			},
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
		&Packet{
			ID: 2,
			DataLink: DataLink{
				SourceMAC:      "00:00:00:00:00:00",
				DestinationMAC: "00:00:00:00:00:00",
				Protocol:       "Ethernet",
			},
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
		&Packet{
			ID: 2,
			DataLink: DataLink{
				SourceMAC:      "00:00:00:00:00:00",
				DestinationMAC: "00:00:00:00:00:00",
				Protocol:       "Ethernet",
			},
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
	})
}

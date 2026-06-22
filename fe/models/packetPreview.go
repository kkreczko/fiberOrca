package models

import "github.com/google/uuid"

// PacketPreview is a struct that represents a preview of a packet
// It implements the Item interface
type PacketPreview struct {
	ID          uuid.UUID
	title       string
	description string
}

func NewPacketPreview(p Packet) PacketPreview {
	return PacketPreview{
		ID:          p.ID,
		title:       p.SourceIP() + ":" + p.DestinationPort(),
		description: p.Datetime().String(),
	}
}

func (p PacketPreview) Title() string {
	return p.title
}

func (p PacketPreview) Description() string {
	return p.description
}

func (p PacketPreview) FilterValue() string {
	return p.Title()
}

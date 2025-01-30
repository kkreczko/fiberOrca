package models

// PacketPreview is a struct that represents a preview of a packet
// It implements the Item interface
type PacketPreview struct {
	ID          int
	title       string
	description string
}

func NewPacketPreview(p Packet) PacketPreview {
	return PacketPreview{
		ID:          p.ID,
		title:       p.SourceIP() + ":" + p.DestinationPort(),
		description: p.TransportProtocol(),
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

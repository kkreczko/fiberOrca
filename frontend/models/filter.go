package models

import "time"
import "github.com/charmbracelet/huh"

var startTime, endTime string

// Filter is a struct that represents a session filter
type Filter struct {
	form                *huh.Form
	IP                  string
	senderPort          string
	receiverPort        string
	networkProtocol     string
	transportProtocol   string
	applicationProtocol string
	startTime           time.Time
	endTime             time.Time
	width               int
	height              int
}

// NewFilter creates a new filter
func NewFilter(width, height int) Filter {
	m := Filter{width: width, height: height}
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Network Protocol").
				Key("Network Protocol").
				Options(huh.NewOptions("IPv4", "IPv6")...).
				Value(&m.networkProtocol),
			huh.NewInput().
				Title("IP").
				Value(&m.IP),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Network Protocol").
				Key("Network Protocol").
				Options(huh.NewOptions("UDP", "TCP")...).
				Value(&m.transportProtocol),
			huh.NewInput().
				Title("Sender Port").
				Placeholder("Sender Port").
				Value(&m.senderPort),
			huh.NewInput().
				Title("Receiver Port").
				Placeholder("Receiver Port").
				Value(&m.receiverPort),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Application Protocol").
				Placeholder("Application Protocol").
				Value(&m.applicationProtocol),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Start Time").
				Placeholder("Start Time").
				Value(&startTime),
			huh.NewInput().
				Title("End Time").
				Placeholder("End Time").
				Value(&endTime),
		),
	)
	return m

}

package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"time"
)

var startTime, endTime string

// Filter is a struct that represents a session filter
type Filter struct {
	form              *huh.Form
	IP                string
	senderPort        string
	receiverPort      string
	transportProtocol string
	startTime         time.Time
	endTime           time.Time
	width             int
	height            int
	active            bool
	session           *Session
}

// NewFilter creates a new filter
func NewFilter(s *Session, width, height int) Filter {
	m := Filter{
		width:   width,
		height:  height,
		active:  false,
		session: s,
	}

	startTime = ""
	endTime = ""

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("IP").
				Value(&m.IP),
			huh.NewSelect[string]().
				Title("Transport Protocol").
				Key("Transport Protocol").
				Options(huh.NewOptions("UDP", "TCP", "UNKNOWN")...).
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
				Title("Start Time").
				Placeholder("YYYY-MM-DD HH:MM:SS").
				Value(&startTime),
			huh.NewInput().
				Title("End Time").
				Placeholder("YYYY-MM-DD HH:MM:SS").
				Value(&endTime),
		),
	).
		WithHeight(height).
		WithWidth(width)

	return m
}

func (m Filter) Init() tea.Cmd {
	return m.form.Init()
}

func (m Filter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m.session, nil
		}
	}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateCompleted {
		// Parse the time strings when form is completed
		if startTime != "" {
			if t, err := time.Parse("2006-01-02 15:04:05", startTime); err == nil {
				m.startTime = t
			}
		}
		if endTime != "" {
			if t, err := time.Parse("2006-01-02 15:04:05", endTime); err == nil {
				m.endTime = t
			}
		}

		// Set filter as active if any field is filled
		m.active = m.IP != "" ||
			m.senderPort != "" ||
			m.receiverPort != "" ||
			m.transportProtocol != "" ||
			!m.startTime.IsZero() ||
			!m.endTime.IsZero()

		// Update session's filter and refresh view
		m.session.filter = m
		m.session.updateFilteredView()

		// Return to session view
		return m.session, nil
	}

	return m, cmd
}

func (m Filter) View() string {
	return m.form.View()
}

// GetTimeRange returns the start and end time of the filter
func (m Filter) GetTimeRange() (time.Time, time.Time) {
	return m.startTime, m.endTime
}

// IsActive returns whether the filter has any active conditions
func (m Filter) IsActive() bool {
	return m.active
}

// Reset clears all filter conditions
func (m *Filter) Reset() {
	m.IP = ""
	m.senderPort = ""
	m.receiverPort = ""
	m.transportProtocol = ""
	m.startTime = time.Time{}
	m.endTime = time.Time{}
	m.active = false

	// Reset the global time variables
	startTime = ""
	endTime = ""

	// Reset the form if it exists
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("IP").
				Value(&m.IP),
			huh.NewSelect[string]().
				Title("Transport Protocol").
				Key("Transport Protocol").
				Options(huh.NewOptions("UDP", "TCP", "UNKNOWN")...).
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
				Title("Start Time").
				Placeholder("YYYY-MM-DD HH:MM:SS").
				Value(&startTime),
			huh.NewInput().
				Title("End Time").
				Placeholder("YYYY-MM-DD HH:MM:SS").
				Value(&endTime),
		),
	).
		WithHeight(height).
		WithWidth(width)
}

// Matches checks if a packet matches the filter criteria
func (m Filter) Matches(packet Packet) bool {
	if !m.active {
		return true
	}

	// Check IP if specified
	if m.IP != "" && packet.SourceIP() != m.IP {
		return false
	}

	// Check ports if specified
	if m.senderPort != "" && packet.SourcePort() != m.senderPort {
		return false
	}
	if m.receiverPort != "" && packet.DestinationPort() != m.receiverPort {
		return false
	}

	// Check protocols if specified
	if m.transportProtocol != "" && packet.Protocol() != m.transportProtocol {
		//fmt.Println(packet.Protocol(), m.transportProtocol)
		return false
	}

	// Check time range if specified
	if !m.startTime.IsZero() && packet.datetime.Before(m.startTime) {
		return false
	}
	if !m.endTime.IsZero() && packet.datetime.After(m.endTime) {
		return false
	}

	return true
}

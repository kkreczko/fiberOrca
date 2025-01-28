package models

import (
	tea "github.com/charmbracelet/bubbletea"
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

func (s *Session) Init() tea.Cmd {
	//TODO implement me
	panic("implement me")
}

func (s *Session) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//TODO implement me
	panic("implement me")
}

func (s *Session) View() string {
	//TODO implement me
	panic("implement me")
}

// NewSession creates a new session
func NewSession() *Session {
	return &Session{}
}

// AddPacket adds a packet to the session
func (s *Session) AddPacket(p Packet) {
	s.collectedPackets = append(s.collectedPackets, p)
}

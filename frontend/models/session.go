package models

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var style = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("241"))

// Session is a struct that represents a application session
// It implements the Model interface
type Session struct {
	collectedPackets list.Model
	filter           Filter
	startTime        time.Time
	endTime          time.Time
	file             string
	ports            []string
	loaded           bool
	quit             bool
}

func (s *Session) Init() tea.Cmd {
	return nil
}

func (s *Session) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !s.loaded {
			style.Width(msg.Width)
			style.Height(msg.Height - 5)
			s.initData(msg.Width, msg.Height-5)
			s.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			s.quit = true
			return s, tea.Quit
		}
	case Packet:
		return s, s.AddPacket(msg)
	}
	var cmd tea.Cmd
	s.collectedPackets, cmd = s.collectedPackets.Update(msg)
	return s, cmd
}

func (s *Session) View() string {
	if s.quit {
		return ""
	}
	if s.loaded {
		return style.Render(s.collectedPackets.View())
	} else {
		return "Loading..."
	}
}

// NewSession creates a new session
func NewSession() *Session {
	return &Session{}
}

// AddPacket adds a packet to the session
func (s *Session) AddPacket(p Packet) tea.Cmd {
	return s.collectedPackets.InsertItem(-1, list.Item(p))
}

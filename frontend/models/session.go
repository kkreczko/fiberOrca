package models

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var modelList = []tea.Model{nil, nil, nil}
var width, height int

const (
	session = iota
	packet
	filter
)

var style = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("241"))

// Session is a struct that represents a application session
// It implements the Model interface
type Session struct {
	packetPreviews   list.Model
	collectedPackets []Packet
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
		width = msg.Width
		height = msg.Height - 5
		if !s.loaded {
			style.Width(width)
			style.Height(height)
			s.initData(width, height)
			s.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			s.quit = true
			return s, tea.Quit
		case "enter":
			ID := s.packetPreviews.SelectedItem().(PacketPreview).ID
			for _, p := range s.collectedPackets {
				if p.ID == ID {
					modelList[packet] = NewPacketView(p, s, width, height)
				}
			}
			modelList[session] = s
			var cmd tea.Cmd
			return modelList[packet].Update(cmd)
		}
	case Packet:
		return s, s.AddPacket(msg)
	}
	var cmd tea.Cmd
	s.packetPreviews, cmd = s.packetPreviews.Update(cmd)
	return s, cmd
}

func (s *Session) View() string {
	if s.quit {
		return ""
	}
	if s.loaded {
		return style.Render(s.packetPreviews.View())
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
	return s.packetPreviews.InsertItem(-1, list.Item(NewPacketPreview(p)))
}

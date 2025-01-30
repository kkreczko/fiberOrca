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
			packetList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
			packetList.Title = "Packets"
			s.packetPreviews = packetList
			style.Width(width)
			style.Height(height)
			s.loaded = true
			s.filter = NewFilter(s, width, height)
		}
	case tea.KeyMsg:
		if !s.loaded {
			return s, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			s.quit = true
			return s, tea.Quit
		case "enter":
			if len(s.packetPreviews.Items()) == 0 {
				return s, nil
			}
			ID := s.packetPreviews.SelectedItem().(PacketPreview).ID
			for _, p := range s.collectedPackets {
				if p.ID == ID {
					modelList[packet] = NewPacketView(p, s, width, height)
				}
			}
			modelList[session] = s
			return modelList[packet].Update(nil)
		case "/":
			// Create a fresh filter and apply it to the session
			newFilter := NewFilter(s, width, height)
			modelList[filter] = newFilter
			return modelList[filter], modelList[filter].Init()
		case "r": // Add a way to reset filter with 'r' key
			s.filter.Reset()
			s.updateFilteredView()
			return s, nil
		}

		var cmd tea.Cmd
		s.packetPreviews, cmd = s.packetPreviews.Update(msg)
		return s, cmd

	case *Packet:
		//fmt.Printf("Received packet: %v\n", msg)
		s.collectedPackets = append(s.collectedPackets, *msg)
		s.updateFilteredView()
		return s, nil
	}

	var cmd tea.Cmd
	s.packetPreviews, cmd = s.packetPreviews.Update(msg)
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
	//packetList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	//packetList.Title = "Packets"

	return &Session{
		startTime: time.Now(),
		loaded:    false,
		quit:      false,
	}
}

func (s *Session) updateFilteredView() {
	var filteredItems []list.Item

	// If filter is not active, show all packets
	if !s.filter.IsActive() {
		for _, packet := range s.collectedPackets {
			filteredItems = append(filteredItems, NewPacketPreview(packet))
		}
	} else {
		// Otherwise, apply filter
		for _, packet := range s.collectedPackets {
			if s.filter.Matches(packet) {
				filteredItems = append(filteredItems, NewPacketPreview(packet))
			}
		}
	}

	s.packetPreviews.SetItems(filteredItems)
}

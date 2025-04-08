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
	PacketPreviews   list.Model
	collectedPackets []Packet
	Filter           *Filter
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
			s.PacketPreviews = packetList
			style.Width(width)
			style.Height(height)
			s.loaded = true
			s.Filter = NewFilter(s, width, height)
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
			if len(s.PacketPreviews.Items()) == 0 {
				return s, nil
			}
			ID := s.PacketPreviews.SelectedItem().(PacketPreview).ID
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
			s.Filter.Reset()
			s.updateFilteredView()
			return s, nil
		}

		var cmd tea.Cmd
		s.PacketPreviews, cmd = s.PacketPreviews.Update(msg)
		return s, cmd

	case *Packet:
		// fmt.Printf("Received packet: %v\n", msg)
		newPacketSlice := append([]Packet{*msg}, s.collectedPackets...)
		s.collectedPackets = newPacketSlice
		s.updateFilteredView()
		return s, nil
	}

	var cmd tea.Cmd
	s.PacketPreviews, cmd = s.PacketPreviews.Update(msg)
	return s, cmd
}

func (s *Session) View() string {
	if s.quit {
		return ""
	}
	if s.loaded {
		return style.Render(s.PacketPreviews.View())
	} else {
		return "Loading..."
	}
}

// NewSession creates a new session
func NewSession() *Session {
	return &Session{
		startTime: time.Now(),
		loaded:    false,
		quit:      false,
		Filter:    nil,
	}
}

func (s *Session) updateFilteredView() {
	var filteredItems []list.Item

	if s.Filter == nil || !s.Filter.IsActive() {
		// If filter is not active, show all packets
		for _, packet := range s.collectedPackets {
			filteredItems = append(filteredItems, NewPacketPreview(packet))
		}
	} else {
		// Apply filter to all packets
		for _, packet := range s.collectedPackets {
			if s.Filter.Matches(packet) {
				filteredItems = append(filteredItems, NewPacketPreview(packet))
			}
		}
	}

	// Set the filtered items
	s.PacketPreviews.SetItems(filteredItems)
}

func (s *Session) GetPackets() []Packet { return s.collectedPackets }

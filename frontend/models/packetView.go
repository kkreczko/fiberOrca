package models

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type PacketView struct {
	packet  Packet
	table   table.Model
	session *Session
}

func (m PacketView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m.session, cmd

		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m PacketView) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func (m PacketView) Init() tea.Cmd {
	return nil
}

func NewPacketView(p Packet, session *Session, width, height int) PacketView {
	width -= 6
	columns := []table.Column{
		{Title: "Key", Width: width / 2},
		{Title: "Data", Width: width / 2},
	}
	rows := []table.Row{
		{"Source MAC", p.SourceMac()},
		{"Destination MAC", p.DestinationMac()},
		{"Protocol", p.DataLinkProtocol()},
		{"Source IP", p.SourceIP()},
		{"Destination IP", p.DestinationIP()},
		{"Protocol", p.NetworkProtocol()},
		{"Source Port", p.SourcePort()},
		{"Destination Port", p.DestinationPort()},
		{"Protocol", p.TransportProtocol()},
		{"TTL", strconv.Itoa(p.TTL)},
		{"Datetime", p.datetime.String()},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithWidth(width),
		table.WithFocused(true),
		table.WithHeight(height),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false).
		Width(width)
	t.SetStyles(s)

	return PacketView{
		packet:  p,
		table:   t,
		session: session,
	}
}

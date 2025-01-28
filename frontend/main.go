package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kruczys/fiberOrca/models"
	"os"
)

var modelList []tea.Model

const (
	session = iota
	filter
)

func main() {
	modelList = []tea.Model{models.NewSession()}
	s := modelList[session]
	p := tea.NewProgram(s)
	if err := p.Start(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}

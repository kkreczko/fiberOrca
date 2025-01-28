package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kruczys/fiberOrca/models"
	"os"
)

func main() {
	s := models.NewSession()
	p := tea.NewProgram(s)
	if err := p.Start(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}

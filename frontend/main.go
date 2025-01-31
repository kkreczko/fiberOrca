package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kruczys/fiberOrca/models"
	"os"
)

func main() {
	ch := make(chan string, 1)
	s := models.NewSession()
	go connectToSocket(s, ch)
	p := tea.NewProgram(s)
	if err := p.Start(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
	ch <- "exit"
	close(ch)
}

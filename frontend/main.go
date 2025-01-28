package main

import (
	tea "github.com/charmbracelet/bubbletea"
	models "github.com/kruczys/fiberOrca/models"
)

var modelList []tea.Model

const (
	sessionIndex = iota
	filterIndex
)

func main() {
	session := models.NewSession()
	modelList[sessionIndex] = session
}

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maitesin/tui-cv/pkg/tabs"
)

func main() {
	tabsHeaders := []string{"Welcome", "Personal", "Skills", "Experience", "Courses", "Looking For"}
	tabContent := []string{"Welcome Tab", "Personal Tab", "Skills Tab", "Experience Tab", "Courses Tab", "Looking For Tab"}
	if _, err := tea.NewProgram(tabs.Model{Tabs: tabsHeaders, TabContent: tabContent}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

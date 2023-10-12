package main

import (
	"fmt"
	"os"

	"github.com/maitesin/tui-cv/pkg/courses"
	"github.com/maitesin/tui-cv/pkg/experience"
	"github.com/maitesin/tui-cv/pkg/looking_for"
	"github.com/maitesin/tui-cv/pkg/skills"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maitesin/tui-cv/pkg/home"
	"github.com/maitesin/tui-cv/pkg/personal"
	"github.com/maitesin/tui-cv/pkg/tabs"
)

func main() {
	tabsHeaders := []string{"Welcome", "Personal", "Skills", "Experience", "Courses", "Looking For"}
	tabContent := []tea.Model{home.Model{}, personal.Model{}, skills.Model{}, experience.Model{}, courses.Model{}, looking_for.Model{}}
	if _, err := tea.NewProgram(tabs.Model{Tabs: tabsHeaders, TabContent: tabContent}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

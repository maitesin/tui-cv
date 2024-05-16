package main

import (
	"fmt"
	"os"

	"github.com/maitesin/tui-cv/pkg/courses"
	"github.com/maitesin/tui-cv/pkg/experience"
	"github.com/maitesin/tui-cv/pkg/looking_for"
	"github.com/maitesin/tui-cv/pkg/skills"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maitesin/tui-cv/pkg/personal"
	"github.com/maitesin/tui-cv/pkg/studies"
	"github.com/maitesin/tui-cv/pkg/tabs"
)

func main() {
	tabsHeaders := []string{"Personal", "Studies", "Skills", "Experience", "Courses", "Looking For"}
	studiesModel, err := studies.NewModel()
	if err != nil {
		panic(err)
	}

	personalModel, err := personal.NewModel()
	if err != nil {
		panic(err)
	}
	tabContent := []tea.Model{personalModel, studiesModel, skills.Model{}, experience.Model{}, courses.Model{}, looking_for.Model{}}
	if _, err := tea.NewProgram(tabs.Model{Tabs: tabsHeaders, TabContent: tabContent}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

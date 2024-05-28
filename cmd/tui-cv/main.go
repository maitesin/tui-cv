package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/maitesin/tui-cv/pkg/experience"
	"github.com/maitesin/tui-cv/pkg/looking_for"
	"github.com/maitesin/tui-cv/pkg/personal"
	"github.com/maitesin/tui-cv/pkg/skills"
	"github.com/maitesin/tui-cv/pkg/studies"
	"github.com/maitesin/tui-cv/pkg/tabs"
	"github.com/maitesin/tui-cv/pkg/welcome"
)

func main() {
	tabsHeaders := []string{"Welcome", "Personal", "Studies", "Skills", "Experience", "Looking For"}

	welcomeModel, err := welcome.NewModel()
	if err != nil {
		panic(err)
	}

	studiesModel, err := studies.NewModel()
	if err != nil {
		panic(err)
	}

	personalModel, err := personal.NewModel()
	if err != nil {
		panic(err)
	}

	skillsModel, err := skills.NewModel()
	if err != nil {
		panic(err)
	}

	experienceModel, err := experience.NewModel()
	if err != nil {
		panic(err)
	}

	lookingForModel, err := looking_for.NewModel()
	if err != nil {
		panic(err)
	}

	tabContent := []tea.Model{welcomeModel, personalModel, studiesModel, skillsModel, experienceModel, lookingForModel}
	if _, err := tea.NewProgram(tabs.Model{Tabs: tabsHeaders, TabContent: tabContent}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

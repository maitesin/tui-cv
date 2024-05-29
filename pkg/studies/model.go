package studies

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `
# University Jaume I (Spain)

- 2010-2013 Bachelor of Engineering in Computer Science
- 2006-2010 Associate Degree in Computer Science

# Courses from Linux Foundation

- August 2019 - LFS258: Kubernetes Fundamentals
- April  2016 - LFD331: Developing Linux Device Drivers
- March  2016 - LFD320: Linux Kernel Internals and Debugging

# Courses from Pragmatic Studio

- April 2023 - Phoenix LiveView Pro v0.20
- March 2022 - Phoenix LiveView Pro v0.15
- July  2021 - Developing with Elixir/OTP

# Courses from PluralSight

- https://app.pluralsight.com/profile/oscar-fornermartinez
`

type Model struct {
	viewport viewport.Model
}

func NewModel() (*Model, error) {
	const width = 72

	vp := viewport.New(width, 24)
	vp.Style = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("62"))

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &Model{
		viewport: vp,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return m.viewport.View()
}

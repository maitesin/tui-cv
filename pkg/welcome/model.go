package welcome

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `
# Welcome to Oscar Forner's interactive CV

Welcome to my interactive CV :)

In this application I will show you some information about myself.

For a more exhaustive into my information, please have a look at my personal website and/or LinkedIn. 

* LinkedIn: https://www.linkedin.com/in/oscarforner
* GitHub: https://github.com/maitesin
* Website: https://oscarforner.com
* Phone: +34 617606012
* Email: oscar.forner.martinez@gmail.com

## Tips
* Use the arrow keys (←|→) to navigate through the tabs.
* Use the "q" key to exit.
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

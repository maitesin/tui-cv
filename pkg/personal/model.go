package personal

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `# Information

Name: Oscar Forner Martinez

Date of Birth: 24/03/1988

Nationality: Spanish

# Languages

English: Fluent

Spanish: Native

Catalan: Native

# About me

I am a Senior Software Engineer interested in Backend Development and Systems Programming.
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

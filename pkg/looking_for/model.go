package looking_for

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `
# What am I looking for?

I am currently looking for new opportunities that would allow me to:

* Design, develop, and maintain high performant reliable systems.
* Create and integrate APIs to expose and extend functionality.
* Create and improve the tools used during the development process.
* Teach and learn about best practices such as DDD, and clean code.
* Contribute to Open Source projects

If you think I can be a good fit for a job opening in your organization do not hesitate on contact me ;)
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

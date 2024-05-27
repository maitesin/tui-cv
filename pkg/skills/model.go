package skills

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const hardContent = `
# Programming Languages

	- Go: 8 years
	- Python: 5 years
	- C++: 4 years
	- Elixir: 3 years

# Orchestration

	- On-premise: 7 years
	- AWS: 5 years
	- Kubernetes: 4 years
	- GCP: 2 years

# Software Engineering

	- TDD, DDD, Clean Code,
	- CQRS, Design Patterns,
	- Microservices
`

type Hard struct {
	viewport viewport.Model
}

func newHard() (*Hard, error) {
	const width = 35

	vp := viewport.New(width, 22)
	vp.Style = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("62"))

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(hardContent)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &Hard{
		viewport: vp,
	}, nil
}

func (m Hard) Init() tea.Cmd {
	return nil
}

func (m Hard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Hard) View() string {
	return m.viewport.View()
}

const softContent = `# CI/CD

- GitHub Actions: 7 years
- Jenkins: 3 years
- Gitlab CI: 2 years

# Database Engines

- PostgreSQL: 8 years
- MongoDB: 4 years

# Soft Skills

- Communication, Self Starter,
- Collaboration, Team Player,
- Problem Solving
`

type Soft struct {
	viewport viewport.Model
}

func newSoft() (*Soft, error) {
	const width = 35

	vp := viewport.New(width, 22)
	vp.Style = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("62"))

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(softContent)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &Soft{
		viewport: vp,
	}, nil
}

func (m Soft) Init() tea.Cmd {
	return nil
}

func (m Soft) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Soft) View() string {
	return m.viewport.View()
}

type Model struct {
	hard *Hard
	soft *Soft
}

func NewModel() (*Model, error) {
	hard, err := newHard()
	if err != nil {
		return nil, err
	}

	soft, err := newSoft()
	if err != nil {
		return nil, err
	}

	return &Model{
		hard: hard,
		soft: soft,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.hard.View(), m.soft.View())
}

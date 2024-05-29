package experience

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const content = `
# 2021 - 2024 Software Engineer at Sketch
IC working with Go and Elixir deploying microservices in AWS 
# 2020 - 2021 Lead Software Engineer at Paack
Leading the BE Go team with multiple microservices in GCP
# 2018 - 2020 Senior Software Engineer at GlobalSign
Most senior member of the Go team deploying on-premise in FreeBSD
# 2016 - 2018 Software Engineer at VCATechnolgy
IC working with C++ and Python to build a device to analyse video
# 2015 - 2016 Software Engineer at PRQA
IC working with C++ and Java to build an static analyser for Java
# 2013 - 2015 Software Engineer at European Bioinformatics Institute
IC working with Python and Java to build processing pipelines
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

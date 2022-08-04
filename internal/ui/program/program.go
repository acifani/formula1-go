package program

import (
	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/results"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	pageResults    = iota
	pageDriverInfo = iota
)

type Page = int

type model struct {
	page    Page
	styles  ui.Styles
	results tea.Model
}

func New(styles ui.Styles) *tea.Program {
	return tea.NewProgram(model{
		page:    pageResults,
		styles:  styles,
		results: results.NewModel(),
	})
}

func (m model) Init() tea.Cmd {
	return m.results.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			cmds = append(cmds, tea.Quit)
		}
	}

	switch m.page {
	case pageResults:
		m.results, cmd = m.results.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.page {
	case pageResults:
		return m.results.View()
	}

	return ""
}

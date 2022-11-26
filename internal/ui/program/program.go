package program

import (
	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/driver"
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/internal/ui/results"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	pageResults    = iota
	pageDriverInfo = iota
)

type Page = int8

type model struct {
	currentPage Page
	pageModels  map[Page]page.Model
	styles      ui.Styles
}

func New(styles ui.Styles) *tea.Program {
	return tea.NewProgram(&model{
		currentPage: pageResults,
		styles:      styles,
	})
}

func (m *model) Init() tea.Cmd {
	m.pageModels = map[Page]page.Model{
		pageResults:    results.New(m.styles),
		pageDriverInfo: driver.New(),
	}

	m.currentPage = pageResults

	return m.getCurrentPageModel().Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			cmds = append(cmds, tea.Quit)
		}
	case driver.DriverLoadedMsg:
		m.currentPage = pageDriverInfo
	}

	currentPageModel := m.getCurrentPageModel()
	if currentPageModel != nil {
		newPageModel, cmd := currentPageModel.Update(msg)
		m.updateCurrentPageModel(newPageModel)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	currentPageModel := m.getCurrentPageModel()
	if currentPageModel != nil {
		return currentPageModel.View()
	}

	return ""
}

func (m *model) getCurrentPageModel() page.Model {
	return m.pageModels[m.currentPage]
}

func (m model) updateCurrentPageModel(newModel page.Model) {
	m.pageModels[m.currentPage] = newModel
}

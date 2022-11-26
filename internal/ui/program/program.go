package program

import (
	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/internal/ui/results"
	"github.com/acifani/formula1-go/internal/ui/wcc"
	"github.com/acifani/formula1-go/internal/ui/wdc"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	pageResults = iota
	pageWCC     = iota
	pageWDC     = iota
)

type Page = int8

type model struct {
	currentPage Page
	pageModels  map[Page]page.Model
	help        help.Model
	keys        keyMap
	styles      ui.Styles
}

func New(styles ui.Styles) *tea.Program {
	return tea.NewProgram(&model{
		currentPage: pageResults,
		help:        help.New(),
		keys:        keys,
		styles:      styles,
	})
}

func (m *model) Init() tea.Cmd {
	m.pageModels = map[Page]page.Model{
		pageResults: results.New(m.styles),
		pageWCC:     wcc.New(m.styles),
		pageWDC:     wdc.New(m.styles),
	}

	m.currentPage = pageResults

	return m.getCurrentPageModel().Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, m.keys.WDC):
			m.currentPage = pageWDC
			cmds = append(cmds, m.getCurrentPageModel().Init())
		case key.Matches(msg, m.keys.WCC):
			m.currentPage = pageWCC
			cmds = append(cmds, m.getCurrentPageModel().Init())
		}
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
	currentPageView := ""
	if currentPageModel != nil {
		currentPageView = currentPageModel.View()
	}

	helpView := m.help.View(m.keys)

	return currentPageView + "\n" + helpView
}

func (m *model) getCurrentPageModel() page.Model {
	return m.pageModels[m.currentPage]
}

func (m model) updateCurrentPageModel(newModel page.Model) {
	m.pageModels[m.currentPage] = newModel
}

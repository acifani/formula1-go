package program

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/internal/ui/quali"
	"github.com/acifani/formula1-go/internal/ui/results"
	"github.com/acifani/formula1-go/internal/ui/season"
	"github.com/acifani/formula1-go/internal/ui/wcc"
	"github.com/acifani/formula1-go/internal/ui/wdc"
)

const (
	PageResults = iota
	PageWCC     = iota
	PageWDC     = iota
	PageSeason  = iota
	PageQuali   = iota
)

type Page = int8

type model struct {
	currentPage Page
	pageModels  map[Page]page.Model
	help        help.Model
	keys        keyMap
	styles      ui.Styles
}

func New(styles ui.Styles, initialPage Page) *tea.Program {
	return tea.NewProgram(&model{
		currentPage: initialPage,
		help:        help.New(),
		keys:        keys,
		styles:      styles,
	})
}

func (m *model) Init() tea.Cmd {
	m.pageModels = map[Page]page.Model{
		PageResults: results.New(m.styles),
		PageWCC:     wcc.New(m.styles),
		PageWDC:     wdc.New(m.styles),
		PageSeason:  season.New(m.styles),
		PageQuali:   quali.New(m.styles),
	}

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
			m.currentPage = PageWDC
			cmds = append(cmds, m.getCurrentPageModel().Init())
		case key.Matches(msg, m.keys.WCC):
			m.currentPage = PageWCC
			cmds = append(cmds, m.getCurrentPageModel().Init())
		case key.Matches(msg, m.keys.Season):
			m.currentPage = PageSeason
			cmds = append(cmds, m.getCurrentPageModel().Init())
		}
	case results.LoadDone:
		m.currentPage = PageResults
	case quali.LoadDone:
		m.currentPage = PageQuali
	case results.BackMsg:
	case quali.BackMsg:
		m.currentPage = PageSeason
		cmds = append(cmds, m.getCurrentPageModel().Init())
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

	pageTitle := "formula1-go"
	if currentPageModel != nil {
		pageTitle = currentPageModel.GetPageTitle()
	}
	return m.styles.Title.Render(pageTitle) +
		m.styles.Base.Render(currentPageView) +
		m.styles.Footer.Render(helpView)
}

func (m *model) getCurrentPageModel() page.Model {
	return m.pageModels[m.currentPage]
}

func (m model) updateCurrentPageModel(newModel page.Model) {
	m.pageModels[m.currentPage] = newModel
}

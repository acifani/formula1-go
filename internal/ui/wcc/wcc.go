package wcc

import (
	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/pkg/api"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	table  table.Model
	styles ui.Styles
	err    error
}

type fetchDone struct {
	err  error
	data *api.ConstructorStandingsTable
}

func New(styles ui.Styles) page.Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Team", Width: 20},
		{Title: "Pts", Width: 4},
		{Title: "Wins", Width: 4},
	}
	t := table.New(table.WithColumns(columns))
	t.SetStyles(styles.Table)

	return &model{table: t, styles: styles}
}

func (m model) GetPageTitle() string {
	return "Constructor Standings"
}

func (m model) Init() tea.Cmd {
	return fetchStandings
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case fetchDone:
		if msg.err != nil {
			m.err = msg.err
		} else {
			rows := generateRows(msg.data)
			m.table.SetHeight(len(rows) + 2)
			m.table.SetRows(rows)
		}
	}

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return m.styles.Paragraph.Render(m.err.Error())
	}

	return m.table.View()
}

func fetchStandings() tea.Msg {
	standings, err := api.GetCurrentConstructorStandings()
	return fetchDone{data: standings, err: err}
}

func generateRows(standings *api.ConstructorStandingsTable) []table.Row {
	rows := make([]table.Row, len(standings.StandingsLists[0].ConstructorStandings))
	for i, standing := range standings.StandingsLists[0].ConstructorStandings {
		rows[i] = table.Row{
			standing.PositionText,
			standing.Constructor.Name,
			standing.Points,
			standing.Wins,
		}
	}
	return rows
}

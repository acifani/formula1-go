package wcc

import (
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/pkg/api"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	table table.Model
}

type fetchDone *api.ConstructorStandingsTable

func New() page.Model {
	columns := []table.Column{
		{Title: "Position", Width: 10},
		{Title: "Constructor", Width: 20},
		{Title: "Points", Width: 10},
		{Title: "Wins", Width: 10},
	}
	table := table.New(table.WithColumns(columns))

	return &model{table: table}
}

func (m model) Init() tea.Cmd {
	return fetchStandings
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case fetchDone:
		rows := generateRows(msg)
		m.table.SetHeight(len(rows))
		m.table.SetRows(rows)
	}

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return m.table.View()
}

func fetchStandings() tea.Msg {
	standings, _ := api.GetCurrentConstructorStandings()
	return fetchDone(standings)
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

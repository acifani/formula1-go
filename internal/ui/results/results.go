package results

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
	data *api.RaceTable
}

func New(styles ui.Styles) page.Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Driver", Width: 25},
		{Title: "Team", Width: 20},
		{Title: "Status", Width: 15},
		{Title: "Time", Width: 12},
		{Title: "Pts", Width: 4},
	}
	table := table.New(table.WithColumns(columns))

	return &model{table: table, styles: styles}
}

func (m model) Init() tea.Cmd {
	return fetchRows
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case fetchDone:
		if msg.err != nil {
			m.err = msg.err
		}
		rows := generateRows(msg.data)
		m.table.SetHeight(len(rows))
		m.table.SetRows(rows)
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

func fetchRows() tea.Msg {
	results, err := api.GetLatestRaceResult()
	return fetchDone{data: results, err: err}
}

func generateRows(results *api.RaceTable) []table.Row {
	rows := make([]table.Row, len(results.Races[0].Results))
	for i, result := range results.Races[0].Results {
		rows[i] = table.Row{
			result.PositionText,
			result.Number + " " + result.Driver.GivenName + " " + result.Driver.FamilyName,
			result.Constructor.Name,
			result.Status,
			result.Time.Time,
			result.Points,
		}
	}
	return rows
}

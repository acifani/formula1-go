package wdc

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/driver"
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/pkg/api"
)

type model struct {
	data   *api.DriverStandingsTable
	table  table.Model
	styles ui.Styles
	err    error
}

type fetchDone struct {
	err  error
	data *api.DriverStandingsTable
}

func New(styles ui.Styles) page.Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Driver", Width: 25},
		{Title: "Team", Width: 20},
		{Title: "Pts", Width: 4},
		{Title: "Wins", Width: 4},
	}
	t := table.New(table.WithColumns(columns), table.WithFocused(true))
	t.SetStyles(styles.SelectableTable)

	return &model{table: t, styles: styles}
}

func (m model) GetPageTitle() string {
	return "Driver Standings"
}

func (m model) Init() tea.Cmd {
	return fetchStandings
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.data != nil {
				idx := m.table.Cursor()
				driverID := m.data.StandingsLists[0].DriverStandings[idx].Driver.DriverID
				return m, driver.LoadResults(m.data.Season, driverID)
			}
		}
	case fetchDone:
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.data = msg.data
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
	standings, err := api.GetCurrentDriverStandings()
	return fetchDone{data: standings, err: err}
}

func generateRows(standings *api.DriverStandingsTable) []table.Row {
	rows := make([]table.Row, len(standings.StandingsLists[0].DriverStandings))
	for i, standing := range standings.StandingsLists[0].DriverStandings {
		rows[i] = table.Row{
			standing.PositionText,
			standing.Driver.PermanentNumber + " " + standing.Driver.GivenName + " " + standing.Driver.FamilyName,
			standing.Constructors[0].Name,
			standing.Points,
			standing.Wins,
		}
	}
	return rows
}

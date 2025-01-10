package results

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/driver"
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/internal/ui/quali"
	"github.com/acifani/formula1-go/pkg/api"
)

type model struct {
	data   *api.RaceTable
	table  table.Model
	styles ui.Styles
	err    error
}

type LoadDone struct {
	err  error
	data *api.RaceTable
}

type BackMsg struct{}

func New(styles ui.Styles) page.Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Driver", Width: 25},
		{Title: "Team", Width: 20},
		{Title: "Status", Width: 15},
		{Title: "Time", Width: 12},
		{Title: "Pts", Width: 4},
	}
	t := table.New(table.WithColumns(columns), table.WithFocused(true))
	t.SetStyles(styles.SelectableTable)

	return &model{table: t, styles: styles}
}

func (m model) GetPageTitle() string {
	if m.data != nil && len(m.data.Races) > 0 {
		return m.data.Races[0].RaceName + " results"
	}
	return "Race results"
}

func (m model) Init() tea.Cmd {
	return fetchRows
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, backMsg
		case "q":
			if m.data != nil {
				return m, quali.LoadResults(m.data.Season, m.data.Round)
			}
		case "enter":
			if m.data != nil {
				idx := m.table.Cursor()
				driverID := m.data.Races[0].Results[idx].Driver.DriverID
				return m, driver.LoadResults(m.data.Season, driverID)
			}
		}
	case LoadDone:
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

func fetchRows() tea.Msg {
	results, err := api.GetLatestRaceResult()
	return LoadDone{data: results, err: err}
}

func LoadResults(year, round string) tea.Cmd {
	return func() tea.Msg {
		results, err := api.GetRaceResult(year, round)
		if len(results.Races) > 0 {
			return LoadDone{data: results, err: err}
		}

		return nil
	}
}

func generateRows(results *api.RaceTable) []table.Row {
	if len(results.Races) == 0 {
		return nil
	}

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

func backMsg() tea.Msg {
	return BackMsg{}
}

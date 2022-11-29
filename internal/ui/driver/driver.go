package driver

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/page"
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
		{Title: "Race", Width: 25},
		{Title: "Pos", Width: 4},
		{Title: "Pts", Width: 4},
		{Title: "Status", Width: 15},
	}
	t := table.New(table.WithColumns(columns))
	t.SetStyles(styles.Table)

	return &model{table: t, styles: styles}
}

func (m model) GetPageTitle() string {
	if m.data != nil {
		driver := m.data.Races[0].Results[0].Driver
		team := m.data.Races[0].Results[0].Constructor
		return driver.PermanentNumber + " " + driver.GivenName + " " + driver.FamilyName + " - " + team.Name
	}
	return ""
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, backMsg
		}
	case LoadDone:
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.data = msg.data
			rows := generateRows(msg.data)
			m.table.SetHeight(len(rows))
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

func LoadResults(year, driverID string) tea.Cmd {
	return func() tea.Msg {
		results, err := api.GetDriverRaceResults(year, driverID)
		return LoadDone{data: results, err: err}
	}
}

func generateRows(data *api.RaceTable) []table.Row {
	rows := make([]table.Row, len(data.Races))
	for i, race := range data.Races {
		rows[i] = table.Row{
			race.Round,
			race.RaceName,
			race.Results[0].PositionText,
			race.Results[0].Points,
			race.Results[0].Status,
		}
	}
	return rows
}

func backMsg() tea.Msg {
	return BackMsg{}
}

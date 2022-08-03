package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"

	"github.com/acifani/formula1-go/pkg/api"
)

type model struct {
	styles Styles
	err    error
	table  table.Model
}

type fetchDone struct {
	err  error
	data *api.RaceTable
}

const (
	columnKeyPosition = "position"
	columnKeyDriver   = "driver"
	columnKeyTeam     = "team"
	columnKeyPoints   = "points"
	columnKeyStatus   = "status"
	columnsKeyTime    = "gap"
)

func NewProgram(styles Styles) *tea.Program {
	return tea.NewProgram(model{
		styles: styles,
	})
}

func (m model) Init() tea.Cmd {
	return fetchRows
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			cmds = append(cmds, tea.Quit)
		}
	case fetchDone:
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.table = table.New([]table.Column{
				table.NewColumn(columnKeyPosition, "#", 4),
				table.NewColumn(columnKeyDriver, "Driver", 20),
				table.NewColumn(columnKeyTeam, "Team", 20),
				table.NewColumn(columnKeyStatus, "Status", 15),
				table.NewColumn(columnsKeyTime, "Time", 12),
				table.NewColumn(columnKeyPoints, "Pts", 4),
			}).WithRows(generateRows(msg.data))
		}
	}

	return m, tea.Batch(cmds...)
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
	race := results.Races[0]
	var rows []table.Row
	for _, result := range race.Results {
		rows = append(rows, table.NewRow(table.RowData{
			columnKeyPosition: result.Position,
			columnKeyDriver:   result.Number + " " + result.Driver.GivenName + " " + result.Driver.FamilyName,
			columnKeyTeam:     result.Constructor.Name,
			columnKeyStatus:   result.Status,
			columnsKeyTime:    result.Time.Time,
			columnKeyPoints:   result.Points,
		}))
	}

	return rows
}

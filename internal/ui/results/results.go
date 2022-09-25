package results

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"

	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/driver"
	"github.com/acifani/formula1-go/pkg/api"
)

type Model struct {
	styles ui.Styles
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
	columnKeyDriverID = "driverId"
)

func New(styles ui.Styles) Model {
	return Model{styles: styles}
}

func (m Model) Init() tea.Cmd {
	return fetchRows
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
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
		case tea.KeyEnter:
			row := m.table.HighlightedRow()
			cmds = append(
				cmds,
				driver.Init(row.Data[columnKeyDriverID].(string)),
			)
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
			}).WithRows(generateRows(msg.data)).Focused(true)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
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

			// Metadata
			columnKeyDriverID: result.Driver.DriverID,
		}))
	}

	return rows
}

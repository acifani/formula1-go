package quali

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/pkg/api"
)

type model struct {
	raceName string
	table    table.Model
	styles   ui.Styles
	err      error
}

type LoadDone struct {
	err  error
	data *api.QualifyingTable
}

type BackMsg struct{}

func New(styles ui.Styles) page.Model {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Driver", Width: 25},
		{Title: "Team", Width: 20},
		{Title: "Q1", Width: 10},
		{Title: "Q2", Width: 10},
		{Title: "Q3", Width: 10},
	}
	t := table.New(table.WithColumns(columns))
	t.SetStyles(styles.Table)

	return &model{table: t, styles: styles}
}

func (m model) GetPageTitle() string {
	return m.raceName + " qualifying results"
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
			m.raceName = msg.data.Races[0].RaceName
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

func LoadResults(year, round string) tea.Cmd {
	return func() tea.Msg {
		results, err := api.GetQualifyingResult(year, round)
		return LoadDone{data: results, err: err}
	}
}

func generateRows(results *api.QualifyingTable) []table.Row {
	rows := make([]table.Row, len(results.Races[0].QualifyingResults))
	for i, result := range results.Races[0].QualifyingResults {
		rows[i] = table.Row{
			result.Position,
			result.Number + " " + result.Driver.GivenName + " " + result.Driver.FamilyName,
			result.Constructor.Name,
			result.Q1,
			result.Q2,
			result.Q3,
		}
	}
	return rows
}

func backMsg() tea.Msg {
	return BackMsg{}
}

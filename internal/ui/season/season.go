package season

import (
	"time"

	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/internal/ui/results"
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
	data *api.ScheduleTable
}

func New(styles ui.Styles) page.Model {
	columns := []table.Column{
		{Title: "Year", Width: 0},
		{Title: "#", Width: 2},
		{Title: "Race", Width: 25},
		{Title: "Circuit", Width: 25},
		{Title: "FP1 Time", Width: 20},
		{Title: "FP2 Time", Width: 20},
		{Title: "FP3 Time", Width: 20},
		{Title: "Quali Time", Width: 20},
		{Title: "Race Time", Width: 20},
	}
	table := table.New(table.WithColumns(columns), table.WithFocused(true))

	return &model{table: table, styles: styles}
}

func (m model) Init() tea.Cmd {
	return fetchSchedule
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			row := m.table.SelectedRow()
			return m, results.LoadResults(row[0], row[1])
		}

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

func fetchSchedule() tea.Msg {
	schedule, err := api.GetCurrentSeasonSchedule()
	return fetchDone{data: schedule, err: err}
}

func generateRows(schedule *api.ScheduleTable) []table.Row {
	rows := make([]table.Row, len(schedule.Races))
	for i, race := range schedule.Races {
		rows[i] = table.Row{
			schedule.Season,
			race.Round,
			race.RaceName,
			race.Circuit.Location.Locality + ", " + race.Circuit.Location.Country,
			formatDate(race.FirstPractice.Date, race.FirstPractice.Time),
			formatDate(race.SecondPractice.Date, race.SecondPractice.Time),
			formatDate(race.ThirdPractice.Date, race.FirstPractice.Time),
			formatDate(race.Qualifying.Date, race.Qualifying.Time),
			formatDate(race.Date, race.Time),
		}
	}
	return rows
}

func formatDate(datePart, timePart string) string {
	d, _ := time.Parse("2006-01-02", datePart)
	t, _ := time.Parse("15:04:05Z", timePart)

	return d.Format("_2 Jan 06") + " " + t.Format("15:04")
}

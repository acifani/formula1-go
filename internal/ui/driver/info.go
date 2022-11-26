package driver

import (
	"github.com/acifani/formula1-go/internal/ui/page"
	"github.com/acifani/formula1-go/pkg/api"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	driver *api.Driver
}

type DriverLoadedMsg *api.Driver

func New() page.Model {
	return &model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case DriverLoadedMsg:
		m.driver = msg
	}

	return m, nil
}

func (m model) View() string {
	return "Name: " + m.driver.GivenName + " " + m.driver.FamilyName +
		"\nNumber: " + m.driver.PermanentNumber + ", Code: " + m.driver.Code +
		"\nBirthdate: " + m.driver.DateOfBirth + ", Nationality: " + m.driver.Nationality
}

func LoadDriver(id string) tea.Cmd {
	return func() tea.Msg {
		driver, _ := api.GetDriverInfo(id)
		return DriverLoadedMsg(driver)
	}
}

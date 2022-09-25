package driver

import (
	"github.com/acifani/formula1-go/pkg/api"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	driver *api.Driver
}

type InitDriverData struct {
	ID string
}

type InitDriverMsg *api.Driver

func New() Model {
	return Model{}
}

func Init(id string) tea.Cmd {
	return func() tea.Msg {
		driver, _ := api.GetDriverInfo(id)
		return InitDriverMsg(driver)
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case InitDriverMsg:
		m.driver = msg
	}

	return m, nil
}

func (m Model) View() string {
	return "Name: " + m.driver.GivenName + m.driver.FamilyName +
		"\nNumber: " + m.driver.PermanentNumber + ", Code: " + m.driver.Code +
		"\nBirthdate: " + m.driver.DateOfBirth + ", Nationality: " + m.driver.Nationality
}

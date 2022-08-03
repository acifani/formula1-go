package ui

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	styles Styles
}

func NewProgram(styles Styles) *tea.Program {
	return tea.NewProgram(model{
		styles: styles,
	})
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return m.styles.Paragraph.Render("Hello, World!")
}

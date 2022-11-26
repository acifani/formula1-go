package season

import (
	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/page"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	styles ui.Styles
}

func New(styles ui.Styles) page.Model {
	return &model{styles: styles}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (page.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return m.styles.Paragraph.Render("Coming soon!")
}

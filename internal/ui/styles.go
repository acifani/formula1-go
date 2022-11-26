package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	indigo       = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	subtleIndigo = lipgloss.AdaptiveColor{Light: "#7D79F6", Dark: "#514DC1"}
	cream        = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
	fuschia      = lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	green        = lipgloss.Color("#04B575")
	red          = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
	faintRed     = lipgloss.AdaptiveColor{Light: "#FF6F91", Dark: "#C74665"}
)

type Styles struct {
	Wrap, Paragraph, Keyword lipgloss.Style
}

func NewStyles() Styles {
	s := Styles{}
	s.Wrap = lipgloss.NewStyle().Width(58)
	s.Paragraph = s.Wrap.Copy().Margin(1, 0, 1, 2)
	s.Keyword = lipgloss.NewStyle().Foreground(green)

	return s
}

package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type Theme struct {
	charColorConfig   charColorConfig
	consoleBackground ConsoleBackground
}

type ConsoleBackground struct {
	consoleColor   lipgloss.ANSIColor
	viewportWidth  int
	viewportHeight int
}

type charColorConfig struct {
	RightStyle     lipgloss.Style
	WrongStyle     lipgloss.Style
	UntouchedStyle lipgloss.Style
}

func CreateCharColorConfig() charColorConfig {
	lipgloss.SetColorProfile(termenv.TrueColor)


	style := charColorConfig{
		RightStyle: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#a6e3a1", Dark: "#a6e3a1"}),
		WrongStyle: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "9", Dark: "9"}),
		UntouchedStyle: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#6c7086", Dark: "#6c7086"}),
	}

	return style
}

func PutBorderAroundText(str string) lipgloss.Style {
	// TODO: Make this width dependent on the viewport width
	return lipgloss.NewStyle().SetString(str).Align(lipgloss.Center).Width(100).Height(30)
}

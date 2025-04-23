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
			Foreground(lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#FFFFFF"}),
		WrongStyle: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "9", Dark: "9"}),
		UntouchedStyle: lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#212124", Dark: "0"}),
	}

	return style
}

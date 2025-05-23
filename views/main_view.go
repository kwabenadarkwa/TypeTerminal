package views

import (
	"fmt"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"

	"github.com/TypeTerminal/theme"
	"github.com/TypeTerminal/utils"
)

type charState int

var (
	displayQuote   utils.Quote
	trackableQuote model
)

const (
	right charState = iota
	wrong
	untouched
)

func TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	displayQuote = getQuote()
	trackableQuote = initialModel()
	return trackableQuote, []tea.ProgramOption{tea.WithAltScreen()}
}

func initialModel() model {
	modelReturn := model{convertQuoteToTrackableType(displayQuote.Quote)}
	// for _, v := range modelReturn.unmarshalledQuotes {
	// 	fmt.Printf("%c %s\n", v.character,v.state)
	// }

	return modelReturn
}

func getQuote() utils.Quote {
	return utils.SelectRandomQuoteFromQuotes(
		utils.GetAllQuotes(filepath.Join("Data", "testWords.json")),
	)
}

func convertQuoteToTrackableType(quote string) []character {
	var arrayThing []character
	for _, v := range quote {
		charThing := character{v, untouched}
		arrayThing = append(arrayThing, charThing)
	}
	return arrayThing
}

type character struct {
	character rune
	state     charState
}

type model struct {
	unmarshalledQuotes []character // items on the to-do list
}

func (m model) Init() tea.Cmd {
	return nil
}

var keyStrokeCount int = 0

func incrementKeyStrokes() {
	keyStrokeCount++
}

func decrementKeyStrokes() {
	if keyStrokeCount > 0 {
		keyStrokeCount--
	}
}

func setPrevCharToUntouched() {
	if keyStrokeCount > 0 {
		trackableQuote.unmarshalledQuotes[keyStrokeCount-1].state = untouched
	}
}

func resetKeyStrokes() {
	keyStrokeCount = 0
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case string(trackableQuote.unmarshalledQuotes[keyStrokeCount].character):
			trackableQuote.unmarshalledQuotes[keyStrokeCount].state = right
			incrementKeyStrokes()

		case "backspace":
			setPrevCharToUntouched()
			decrementKeyStrokes()

		default:
			trackableQuote.unmarshalledQuotes[keyStrokeCount].state = wrong
			incrementKeyStrokes()
		}
		if len(trackableQuote.unmarshalledQuotes) == keyStrokeCount {
			resetKeyStrokes()
		}
	}
	return m, nil
}

func (m model) View() string {
	style := theme.CreateCharColorConfig()
	s := ""
	for _, v := range trackableQuote.unmarshalledQuotes {
		switch v.state {
		case untouched:
			s += fmt.Sprint(style.UntouchedStyle.Render(string(v.character)))
		case right:
			s += fmt.Sprint(style.RightStyle.Render(string(v.character)))
		case wrong:
			s += fmt.Sprint(style.WrongStyle.Render(string(v.character)))
		}
	}
	return s
}

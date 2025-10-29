package views

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"

	"github.com/TypeTerminal/theme"
	"github.com/TypeTerminal/utils"
)

var (
	displayQuote utils.Quote
	sessionState model
)

const (
	right charState = iota
	wrong
	untouched
)

type charState int

type character struct {
	character rune
	state     charState
}

// TODO: should probably be an atomic bool to make it thread safe or whatever
type model struct {
	unmarshalledQuote []character // items on the to-do list
	wpm               int
	wpmTracked        bool
	startTime         time.Time
	endTime           time.Time
}

func TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	resetKeyStrokes()
	displayQuote = getQuote()
	sessionState = initialModel()
	return sessionState, []tea.ProgramOption{tea.WithAltScreen()}
}

func initialModel() model {
	modelReturn := model{
		unmarshalledQuote: convertQuoteToTrackableType(displayQuote.Quote),
		wpm:               0,
		wpmTracked:        false,
	}
	return modelReturn
}

// TODO: this quote getting would be replaced by some API
func getQuote() utils.Quote {
	return utils.SelectRandomQuoteFromQuotes(
		utils.GetAllQuotes(filepath.Join("Data", "testWords.json")),
	)
}

func convertQuoteToTrackableType(quote string) []character {
	var charArray []character
	for _, v := range quote {
		char := character{v, untouched}
		charArray = append(charArray, char)
	}
	return charArray
}

var keyStrokeCount int = 0

func incrementKeyStrokes() {
	wordLen := len(sessionState.unmarshalledQuote)
	if keyStrokeCount < wordLen {
		keyStrokeCount++
	}
}

func decrementKeyStrokes() {
	if keyStrokeCount > 0 {
		keyStrokeCount--
	}
}

func setPrevCharToUntouched() {
	if keyStrokeCount > 0 {
		sessionState.unmarshalledQuote[keyStrokeCount-1].state = untouched
	}
}

func resetKeyStrokes() {
	keyStrokeCount = 0
}

func (m model) setWPM() {
	elapsedTime := m.endTime.Sub(m.startTime)
	log.Println("This is the elapsed time", elapsedTime.Minutes())
}

func (m model) Init() tea.Cmd {
	return nil
}

// TODO: something smells wrong here with the way the code is, I should probably consider refactoring it
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.wpmTracked {
		m.startTime = time.Now()
		m.wpmTracked = true
	}

	if len(sessionState.unmarshalledQuote) == keyStrokeCount {
		m.endTime = time.Now()
		m.setWPM()
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "backspace":
				setPrevCharToUntouched()
				decrementKeyStrokes()
			}
		}

		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case string(sessionState.unmarshalledQuote[keyStrokeCount].character):
			sessionState.unmarshalledQuote[keyStrokeCount].state = right
			incrementKeyStrokes()

		case "backspace":
			setPrevCharToUntouched()
			decrementKeyStrokes()

		default:
			sessionState.unmarshalledQuote[keyStrokeCount].state = wrong
			incrementKeyStrokes()
		}
	}

	return m, nil
}

func (m model) View() string {
	style := theme.CreateCharColorConfig()
	s := ""
	for _, v := range sessionState.unmarshalledQuote {
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

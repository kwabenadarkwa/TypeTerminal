package views

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
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
	unmarshalledQuote []character
	wordCount         int
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
		unmarshalledQuote: unmarshallQuoteToChar(displayQuote.Quote),
		wordCount:         getWordCount(displayQuote.Quote),
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

func getWordCount(quote string) int {
	return len(strings.Split(quote, " "))
}

func unmarshallQuoteToChar(quote string) []character {
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
	elapsedTime := m.endTime.Sub(m.startTime).Seconds()
	log.Println("This is the elapsed time", elapsedTime)
	// this is the naive words per minute not accounting for the fact that some of the words were wrong
	log.Println("Total word count", m.wordCount)
	m.wpm = int((m.wordCount * 60) / int(elapsedTime))
	log.Println("This is the wpm", m.wpm)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.wpmTracked {
		m.startTime = time.Now()
		m.wpmTracked = true
	}

	if len(m.unmarshalledQuote) == keyStrokeCount {
		m.endTime = time.Now()
		m.setWPM()
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			if keyMsg.String() == "backspace" {
				setPrevCharToUntouched()
				decrementKeyStrokes()
			}
		}
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case string(m.unmarshalledQuote[keyStrokeCount].character):
			m.unmarshalledQuote[keyStrokeCount].state = right
			incrementKeyStrokes()

		case "backspace":
			setPrevCharToUntouched()
			decrementKeyStrokes()

		default:
			m.unmarshalledQuote[keyStrokeCount].state = wrong
			incrementKeyStrokes()
		}
	}

	return m, nil
}

func (m model) View() string {
	style := theme.CreateCharColorConfig()
	s := ""
	for _, v := range m.unmarshalledQuote {
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

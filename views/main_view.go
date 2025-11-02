package views

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	charLength        int
	consoleWidth      int
	consoleHeight     int
	wpm               int
	wpmTracked        bool
	accuracy          int
	typingDone        bool
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
		accuracy:          0,
		wpmTracked:        false,
		typingDone:        false,
	}
	charCount := len(modelReturn.unmarshalledQuote)
	modelReturn.charLength = charCount

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

func (m *model) setWPM() {
	elapsedTime := m.endTime.Sub(m.startTime).Seconds()
	log.Println("elapsed time", elapsedTime)

	var wrongCount int
	for _, char := range m.unmarshalledQuote {
		if char.state == wrong {
			wrongCount++
		}
	}
	log.Println("this is the total char length", m.charLength)

	wpmWithoutMistakes := int((m.wordCount * 60) / int(elapsedTime))
	log.Println("Wrong Count", wrongCount)
	wrongCharPercentage := float64(wrongCount) / float64(m.charLength)
	log.Println("WPM regularly without mistakes", wpmWithoutMistakes)
	log.Println("Percentage of mistakes", wrongCharPercentage)
	log.Println("difference with one", 1-wrongCharPercentage)
	m.accuracy = int((float64(1) - wrongCharPercentage) * 100)
	m.wpm = int(float64(wpmWithoutMistakes) * (float64(1) - wrongCharPercentage))
	log.Println("Actual WPM with mistakes", m.wpm)
	log.Println("This is the accuracy", m.accuracy)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.wpmTracked {
		m.startTime = time.Now()
		m.wpmTracked = true
	}

	if len(m.unmarshalledQuote) != keyStrokeCount {
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
		case tea.WindowSizeMsg:
			m.consoleHeight = msg.Height
			m.consoleWidth = msg.Width
		}
	}

	if len(m.unmarshalledQuote) == keyStrokeCount {
		// INFO: this is done this way currently and that might not be the best decision but I don't recalculate
		// the typing speed when the user is done and then backspaces because I make the assumption that
		// they wouldn't necessarily want to redo their entire text
		// TODO: I should probably think of having a redo feature of some kind
		if !m.typingDone {
			m.endTime = time.Now()
			m.setWPM()
			m.typingDone = true
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "backspace":
				setPrevCharToUntouched()
				decrementKeyStrokes()
			}
		case tea.WindowSizeMsg:
			m.consoleHeight = msg.Height
			m.consoleWidth = msg.Width
		}

	}

	return m, nil
}

// TODO:  Complete Adaptive Colors use this to define the light and dark theme colors
func (m model) View() string {
	style := theme.CreateCharColorConfig()

	s := ""
	applicationName := lipgloss.NewStyle().
		SetString("typeTerm").
		Align(lipgloss.Left).
		Foreground(lipgloss.Color("#f08c00")).
		Width(m.consoleWidth / 6).
		String()

	// TODO: I need another thing to track which page we are currently on
	// and I need to separate the things out into their various parts
	instruction := lipgloss.NewStyle().
		SetString("practice • stats").
		Align(lipgloss.Right).
		Width(m.consoleWidth / 3).
		String()

	header := lipgloss.NewStyle().
		SetString(lipgloss.JoinHorizontal(lipgloss.Center, applicationName, instruction)).
		Border(lipgloss.NormalBorder()).
		String()

	wpm := lipgloss.NewStyle().
		SetString("WPM: ", strconv.Itoa(m.wpm)).
		Width(m.consoleWidth/10).
		Align(lipgloss.Left).
		String()

	accuracy := lipgloss.NewStyle().
		Width(m.consoleWidth/7).
		SetString("Accuracy: ", strconv.Itoa(m.accuracy), "%").
		Align(lipgloss.Right).String()

	currentStats := lipgloss.NewStyle().
		SetString(lipgloss.JoinHorizontal(lipgloss.Center, wpm, accuracy)).
		Width(m.consoleWidth / 4).
		Border(lipgloss.HiddenBorder()).
		String()

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

	mainContent := lipgloss.NewStyle().
		BorderStyle(lipgloss.HiddenBorder()).
		SetString(s).
		Width(m.consoleWidth / 2).
		Height(5).
		Align(lipgloss.Center).
		String()

	footer := lipgloss.NewStyle().
		SetString("tab: new text   ⬅️➡️ : navigation   ctrl+c: quit").
		BorderStyle(lipgloss.HiddenBorder()).
		Align(lipgloss.Center).
		Width(m.consoleWidth / 2).
		Foreground(lipgloss.Color("#6c7086")).
		String()

	stackedContent := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		currentStats,
		mainContent,
		footer,
	)

	return lipgloss.Place(
		m.consoleWidth,
		m.consoleHeight,
		lipgloss.Center,
		lipgloss.Center,
		stackedContent,
	)
}

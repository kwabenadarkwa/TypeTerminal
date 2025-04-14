package screens

import (
	"fmt"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"

	utils "github.com/TypeTerminal/Utils"
)

type charState int

var (
	displayQuote   utils.Quote
	trackableQuote model
)

var (
	rightStyle     lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	wrongStyle     lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	untouchedStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("0"))
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
	// TODO: figure out a better name for this section
	// TODO: there might be more things that are needed in here
	unmarshalledQuotes []character // items on the to-do list
	// TODO: this should keep track of the index that the user is currently on
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}
	// switch msg := msg.(type) {
	// // Is it a key press?
	// case tea.KeyMsg:

	//
	// 	// Cool, what was the actual key pressed?
	// 	switch msg.String() {
	//
	// 	// These keys should exit the program.
	// 	case "ctrl+c", "q":
	// 		return m, tea.Quit
	//
	// 	// The "up" and "k" keys move the cursor up
	// 	case "up", "k":
	// 		if m.cursor > 0 {
	// 			m.cursor--
	// 		}
	//
	// 	// The "down" and "j" keys move the cursor down
	// 	case "down", "j":
	// 		if m.cursor < len(m.choices)-1 {
	// 			m.cursor++
	// 		}
	//
	// 	// The "enter" key and the spacebar (a literal space) toggle
	// 	// the selected state for the item that the cursor is pointing at.
	// 	case "enter", " ":
	// 		_, ok := m.selected[m.cursor]
	// 		if ok {
	// 			delete(m.selected, m.cursor)
	// 		} else {
	// 			m.selected[m.cursor] = struct{}{}
	// 		}
	// 	}
	// }
	//
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	s := ""
	for _, v := range trackableQuote.unmarshalledQuotes {
		switch v.state {
		case untouched:
			s += fmt.Sprint(untouchedStyle.Render(string(v.character)))
		case right:
			s += fmt.Sprint(rightStyle.Render(string(v.character)))
		case wrong:
			s += fmt.Sprint(wrongStyle.Render(string(v.character)))
		}
	}
	return s
}

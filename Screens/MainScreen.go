package screens

import (
	"fmt"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"

	utils "github.com/TypeTerminal/Utils"
)

var displayQuote utils.Quote

func TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	displayQuote = getQuote()
	return initialModel(), []tea.ProgramOption{tea.WithAltScreen()}
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

// func updateShowingQuote() {
// }

func convertQuoteToTrackableType(quote string) []character {
	var arrayThing []character
	for _, v := range quote {
		// TODO: make this an enum
		charThing := character{v, "untouched"}
		arrayThing = append(arrayThing, charThing)

	}

	return arrayThing
}

// this is what shows whether a particular character has been pressed or not and helps us keep track of each
// thing and whether it has been pressed or not
type character struct {
	character rune
	state     string
}

type model struct {
	// TODO: figure out a better name for this section
	// TODO: there might be more things that are needed in here
	unmarshalledQuotes []character // items on the to-do list
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	s := displayQuote.Quote + "\n"

	// // Iterate over our choices
	// for i, choice := range m.choices {
	//
	// 	// Is the cursor pointing at this choice?
	// 	cursor := " " // no cursor
	// 	if m.cursor == i {
	// 		cursor = ">" // cursor!
	// 	}
	//
	// 	// Is this choice selected?
	// 	checked := " " // not selected
	// 	if _, ok := m.selected[i]; ok {
	// 		checked = "x" // selected!
	// 	}
	//
	// 	// Render the row
	// 	s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	// }
	//
	// // The footer
	// s += "\nPress q to quit.\n"
	//
	// // Send the UI for rendering

	return s
}

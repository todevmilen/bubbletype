package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// styles
var (
	correctStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	defaultStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
)

type model struct {
	sentence      string
	position      int
	inputError    bool
	width, height int
	startTime     time.Time
	endTime       time.Time
	hasFinished   bool
}

func initialModel(height, width int) model {
	return model{
		sentence:   "He liked to play with words in the bathtub. Joyce enjoyed eating pancakes with ketchup.",
		position:   0,
		inputError: false,
		width:      width,
		height:     height,
	}
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
			// TODO: restart the program
		case "ctrl+r":
			return initialModel(m.height, m.width), nil
		}
		if m.position == 1 {
			m.startTime = time.Now()
		}
		if m.hasFinished {
			return m, nil
		}
		if m.position < len(m.sentence) {
			expectedChar := string(m.sentence[m.position])

			if msg.String() == expectedChar {
				m.position++
				m.inputError = false

				if m.position == len(m.sentence) {
					m.endTime = time.Now()
					m.hasFinished = true
				}
			} else {
				m.inputError = true
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}
func (m model) View() string {
	var builder strings.Builder
	if m.hasFinished {
		// Calculate total time in seconds
		totalTime := m.endTime.Sub(m.startTime).Seconds()

		// Calculate WPM (words per minute)
		totalChars := len(m.sentence)
		wordsTyped := float64(totalChars) / 5
		wpm := (wordsTyped / totalTime) * 60 // WPM formula

		return fmt.Sprintf("Congrats! You've typed the sentence.\nYour WPM: %.2f", wpm)
	}

	// Build the sentence view
	for i, char := range m.sentence {
		if i < m.position {
			// Highlight correct characters
			builder.WriteString(correctStyle.Render(string(char)))
		} else if i == m.position && m.inputError {
			// Highlight the expected character in red when input is incorrect
			builder.WriteString(errorStyle.Render(string(char)))
		} else {
			// Show untyped characters in white
			builder.WriteString(defaultStyle.Render(string(char)))
		}
	}

	// Additional hint message when there's an error
	if m.inputError {
		builder.WriteString("\nIncorrect key! Try again.")
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, builder.String())
}

func main() {
	m := model{
		sentence: "Hello",
		// sentence: "He liked to play with words in the bathtub. Joyce enjoyed eating pancakes with ketchup.",
		position: 0,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

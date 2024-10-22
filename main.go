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

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			Align(lipgloss.Center).
			Padding(1, 2)

	wpmStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42")) // Green for WPM

	accStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")) // Blue for Accuracy

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 1).
			Margin(1, 0).
			Align(lipgloss.Center)
)

type model struct {
	sentence        string
	position        int
	inputError      bool
	width, height   int
	startTime       time.Time
	endTime         time.Time
	hasFinished     bool
	inputErrorCount int
	hasMistake      bool
}

func initialModel(height, width int) model {
	return model{
		sentence:        "He liked to play with words.",
		position:        0,
		inputError:      false,
		width:           width,
		height:          height,
		inputErrorCount: 0,
		hasMistake:      false,
	}
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("BubbleType")
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "q", "esc", "ctrl+c":
			return m, tea.Quit

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
				m.hasMistake = false

				if m.position == len(m.sentence) {
					m.endTime = time.Now()
					m.hasFinished = true
				}

			} else if !m.hasMistake {
				m.inputErrorCount++
				m.inputError = true
				m.hasMistake = true
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
		wordsTyped := float64(totalChars / 5)
		wpm := (wordsTyped / totalTime) * 60 // WPM formula

		// Calculate accuracy
		correctChars := totalChars - m.inputErrorCount
		acc := int((float64(correctChars) / float64(totalChars)) * 100)

		// return fmt.Sprintf("Congrats! You've typed the sentence.\nYour WPM: %.0f and Your Acc: %d%%", wpm, acc)
		// Build the results screen
		title := titleStyle.Render("🎉 Congrats! 🎉\nYou've completed the sentence.")

		wpmDisplay := wpmStyle.Render(fmt.Sprintf("WPM: %.0f", wpm))
		accDisplay := accStyle.Render(fmt.Sprintf("Accuracy: %d%%", acc))

		// Wrap WPM and accuracy in a border
		wpmBox := borderStyle.Render(wpmDisplay)
		accBox := borderStyle.Render(accDisplay)

		// Combine everything into the final result view
		resultView := lipgloss.JoinVertical(lipgloss.Center, title, wpmBox, accBox)

		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, resultView)

	}

	// Build the sentence view
	for i, char := range m.sentence {
		if i < m.position {
			// Highlight correct characters with green
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
	m := initialModel(0, 0)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

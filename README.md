# Bubble Type - Typing Speed Game

A simple terminal-based typing game built using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) TUI framework in Go. Test your typing speed and accuracy by typing out sentences and get your words-per-minute (WPM) score at the end!

## Features

- Displays a sentence for the user to type.
- Highlights correct characters in green as you type.
- Shows incorrect characters in red.
- Calculates and displays words per minute (WPM) after completing the sentence.

## Demo

![Typing Speed Game Demo](demo.gif)

## Prerequisites

Make sure you have Go installed on your machine. You can download Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/todevmilen/bubbletype.git
   ```

2. Navigate to the project directory:

   ```bash
   cd bubbletype
   ```

3. Install the dependencies (Bubble Tea and Lip Gloss):

   ```bash
   go get github.com/charmbracelet/bubbletea
   go get github.com/charmbracelet/lipgloss
   ```

## How to Run

1. Build the game:

   ```bash
   go build
   ```

2. Run the game:

   ```bash
   ./bubbletype
   ```

## How to Play

- When the game starts, a sentence will appear in the terminal.
- Type the sentence as quickly and accurately as possible.
- Correct characters will be highlighted in green.
- Incorrect characters will be shown in red.
- When you finish typing the sentence, your Words Per Minute (WPM) score will be displayed.

## WPM Calculation

WPM (Words Per Minute) is calculated using the formula:

\[
\text{WPM} = \left( \frac{\text{characters typed}}{5} \right) \times \left( \frac{60}{\text{seconds elapsed}} \right)
\]

The average word length is considered to be 5 characters, including spaces and punctuation.

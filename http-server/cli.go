package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI struct to handle the command line application
type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

// PlayPoker start a poker game
func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.in)
	reader.Scan()
	c.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

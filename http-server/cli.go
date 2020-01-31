package poker

import "io"

// CLI struct to handle the command line application
type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

// PlayPoker start a poker game
func (c *CLI) PlayPoker() {
	c.playerStore.RecordWin("Chris")
}

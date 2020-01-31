package poker

// CLI struct to handle the command line application
type CLI struct {
	playerStore PlayerStore
}

// PlayPoker start a poker game
func (c *CLI) PlayPoker() {
	c.playerStore.RecordWin("Cleo")
}

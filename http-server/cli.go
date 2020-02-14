package poker

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

// PlayerPrompt message to ask how many players the game will have
const PlayerPrompt = "Please enter the number of players: "

// CLI struct to handle the command line application
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	out         io.Writer
	alerter     BlindAlerter
}

// NewCLI return a new CLI struct based on a given store and io.Reader
func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		store,
		bufio.NewScanner(in),
		out,
		alerter,
	}
}

// PlayPoker start a poker game
func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	c.scheduleBlindAlerts()
	userInput := c.readLine()
	c.playerStore.RecordWin(extractWinner(userInput))
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func (c *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, b := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, b)
		blindTime = blindTime + 10*time.Minute
	}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

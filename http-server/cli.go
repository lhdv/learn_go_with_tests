package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// PlayerPrompt message to ask how many players the game will have
const PlayerPrompt = "Please enter the number of players: "

// BadPlayerInputErrMsg message when a non-numeric value for number of players is entered
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

// CLI struct to handle the command line application
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// NewCLI return a new CLI struct based on a given store and io.Reader/Writer
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

// PlayPoker start a poker game
func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayersInput := c.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}

	c.game.Start(numberOfPlayers, c.out)

	winnerInput := c.readLine()
	winner := extractWinner(winnerInput)

	c.game.Finish(winner)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

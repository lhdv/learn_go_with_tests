package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/lhdv/learn_go_with_tests/http-server"
)

type GameSpy struct {
	StartedWith int
	FinishWith  string
}

func (g *GameSpy) Start(numberOfPlayer int) {
	g.StartedWith = numberOfPlayer
}

func (g *GameSpy) Finish(winner string) {
	g.FinishWith = winner
}

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it prompts the user to enter the number of players and start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}

	})

}

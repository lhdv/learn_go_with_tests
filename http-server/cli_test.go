package poker_test

import (
	"bytes"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	poker "github.com/lhdv/learn_go_with_tests/http-server"
)

type GameSpy struct {
	StartedCalled bool
	StartedWith   int
	FinishWith    string
}

func (g *GameSpy) Start(numberOfPlayer int, alerterDest io.Writer) {
	g.StartedCalled = true
	g.StartedWith = numberOfPlayer
}

func (g *GameSpy) Finish(winner string) {
	g.FinishWith = winner
}

var (
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
	dummyGame         = &GameSpy{}
)

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

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}

	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartedCalled {
			t.Errorf("game should not have started")
		}

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}

		out := &bytes.Buffer{}
		in := strings.NewReader("3\nChris wins")

		poker.NewCLI(in, out, game).PlayPoker()

		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")

		if game.StartedCalled {
			t.Errorf("game should not have started")
		}
	})

	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		game := &GameSpy{}
		winner := "Ruth"

		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		defer server.Close()

		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
	})
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, want)
	}
}

func assertGameStartedWith(t *testing.T, game *GameSpy, players int) {
	t.Helper()

	if game.StartedWith != players {
		t.Errorf("wanted Start called with %d but got %d", players, game.StartedWith)
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	if game.FinishWith != winner {
		t.Errorf("wanted Finish called with %v but got %q", winner, game.FinishWith)
	}
}

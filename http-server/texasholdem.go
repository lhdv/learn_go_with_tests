package poker

import (
	"io"
	"time"
)

// TexasHoldem is an interface to communicate between CLI and the poker game
type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

// NewTexasHoldem create a new TexasHoldem pointer
func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter,
		store,
	}
}

// Start a game given a number of players and set blinds alerter
func (p *TexasHoldem) Start(numberOfPlayers int, alerterDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	// blindIncrement := time.Duration(5+numberOfPlayers) * time.Second

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, b := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, b, alerterDestination)
		blindTime = blindTime + blindIncrement
	}
}

//Finish a game and set a winner
func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}

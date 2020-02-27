package poker

import (
	"fmt"
	"io"
	"testing"
	"time"
)

// StubPlayerStore struct for testing purpouses (implementation of PlayerStore interface)
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

// GetPlayerScore from StubPlayerStore
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

// RecordWin in StubPlayerStore
func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

// GetLeague from StubPlayerStore
func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

// AssertPlayerWin check a winner from StubPlayerStore
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}

// ScheduledAlert keeps a duration of a given blind value will last
type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

// SpyBlindAlerter for testing purpouses
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

// ScheduleAlertAt schedule an alert for a specific duration
func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

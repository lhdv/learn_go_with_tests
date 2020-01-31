package poker

import "sync"

// InMemoryPlayerStore is a temp storage in-memory for player's score
type InMemoryPlayerStore struct {
	mu    sync.RWMutex
	store map[string]int
}

// NewInMemoryPlayerStore returns a new in-memory storage
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		sync.RWMutex{},
		map[string]int{},
	}
}

// GetPlayerScore will get player's score from memory
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.RLock()
	defer i.mu.RUnlock()
	score := i.store[name]
	return score
}

// RecordWin will set player's score
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

// GetLeague gets all player's score from memory
func (i *InMemoryPlayerStore) GetLeague() League {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

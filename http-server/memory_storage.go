package main

// InMemoryPlayerStore is a temp storage in-memory for player's score
type InMemoryPlayerStore struct {
	store map[string]int
}

// NewInMemoryPlayerStore returns a new in-memory storage
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// GetPlayerScore will get player's score from memory
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	score := i.store[name]
	return score
}

// RecordWin will set player's score
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

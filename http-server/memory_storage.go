package main

import "sync"

// InMemoryPlayerStore is a temp storage in-memory for player's score
type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

// NewInMemoryPlayerStore returns a new in-memory storage
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		sync.Mutex{},
		map[string]int{},
	}
}

// GetPlayerScore will get player's score from memory
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	score := i.store[name]
	return score
}

// RecordWin will set player's score
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

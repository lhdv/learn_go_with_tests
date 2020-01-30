package main

import (
	"encoding/json"
	"os"
	"sync"
)

// FileSystemPlayerStore save player data in file system
type FileSystemPlayerStore struct {
	mu       sync.RWMutex
	database *json.Encoder
	league   League
}

// NewFileSystemPlayerStore create a new store
func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)
	return &FileSystemPlayerStore{
		sync.RWMutex{},
		json.NewEncoder(&tape{file}),
		league,
	}
}

// GetLeague return a Player array
func (fs *FileSystemPlayerStore) GetLeague() League {
	return fs.league
}

// GetPlayerScore return a player's score by its name
func (fs *FileSystemPlayerStore) GetPlayerScore(name string) int {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	player := fs.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin increment player's score
func (fs *FileSystemPlayerStore) RecordWin(name string) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	player := fs.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		fs.league = append(fs.league, Player{name, 1})
	}

	fs.database.Encode(fs.league)
}

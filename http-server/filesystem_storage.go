package main

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

// FileSystemPlayerStore save player data in file system
type FileSystemPlayerStore struct {
	mu       sync.RWMutex
	database io.Writer
	league   League
}

// NewFileSystemPlayerStore create a new store
func NewFileSystemPlayerStore(db *os.File) *FileSystemPlayerStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)
	return &FileSystemPlayerStore{
		sync.RWMutex{},
		&tape{db},
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

	json.NewEncoder(fs.database).Encode(fs.league)
}

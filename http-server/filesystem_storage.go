package main

import (
	"encoding/json"
	"io"
	"sync"
)

// FileSystemPlayerStore save player data in file system
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	mu       sync.RWMutex
}

// NewFileSystemPlayerStore create a new store
func NewFileSystemPlayerStore(db io.ReadWriteSeeker) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{
		db,
		sync.RWMutex{},
	}
}

// GetLeague return a Player array
func (fs *FileSystemPlayerStore) GetLeague() League {
	fs.database.Seek(0, 0)
	league, _ := NewLeague(fs.database)
	return league
}

// GetPlayerScore return a player's score by its name
func (fs *FileSystemPlayerStore) GetPlayerScore(name string) int {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	player := fs.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin increment player's score
func (fs *FileSystemPlayerStore) RecordWin(name string) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	league := fs.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	fs.database.Seek(0, 0)
	json.NewEncoder(fs.database).Encode(league)
}

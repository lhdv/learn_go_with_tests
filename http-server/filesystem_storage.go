package main

import (
	"io"
)

// FileSystemPlayerStore save player data in file system
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague return a Player array
func (fs *FileSystemPlayerStore) GetLeague() []Player {
	fs.database.Seek(0, 0)
	league, _ := NewLeague(fs.database)
	return league
}

// GetPlayerScore return a player's score by its name
func (fs *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range fs.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}

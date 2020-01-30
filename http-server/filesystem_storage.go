package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore save player data in file system
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
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

// RecordWin increment player's score
func (fs *FileSystemPlayerStore) RecordWin(name string) {
	league := fs.GetLeague()

	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}

	fs.database.Seek(0, 0)
	json.NewEncoder(fs.database).Encode(league)
}

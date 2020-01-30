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
func (fs *FileSystemPlayerStore) GetLeague() League {
	fs.database.Seek(0, 0)
	league, _ := NewLeague(fs.database)
	return league
}

// GetPlayerScore return a player's score by its name
func (fs *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := fs.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin increment player's score
func (fs *FileSystemPlayerStore) RecordWin(name string) {
	league := fs.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}

	fs.database.Seek(0, 0)
	json.NewEncoder(fs.database).Encode(league)
}

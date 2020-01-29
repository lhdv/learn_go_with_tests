package main

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore save player data in file system
type FileSystemPlayerStore struct {
	database io.Reader
}

// GetLeague return a Player array
func (fs *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player

	json.NewDecoder(fs.database).Decode(&league)

	return league
}

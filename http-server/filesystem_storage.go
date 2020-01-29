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

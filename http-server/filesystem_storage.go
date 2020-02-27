package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"
)

// FileSystemPlayerStore save player data in file system
type FileSystemPlayerStore struct {
	mu       sync.RWMutex
	database *json.Encoder
	league   League
}

// FileSystemPlayerStoreFromFile return a new FileSystemPlayerStore where
// its database is stored in *path* and return a close function
func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v", err)
	}

	return store, closeFunc, nil
}

// NewFileSystemPlayerStore create a new store
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := initPlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		sync.RWMutex{},
		json.NewEncoder(&Tape{file}),
		league,
	}, nil
}

// GetLeague return a Player array
func (fs *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(fs.league, func(i, j int) bool {
		return fs.league[i].Wins > fs.league[j].Wins
	})
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

func initPlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte(`[]`))
		file.Seek(0, 0)
	}

	return nil
}

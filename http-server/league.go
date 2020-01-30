package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// League is a set of players and its scores
type League []Player

// NewLeague expects to decode a json with player's score
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player

	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

// Find will return a player from the league
func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}

	return nil
}

package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer handle all functions of a player
type PlayerServer struct {
	store PlayerStore
}

// ServeHTTP serve to handle GET/POST requests
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	p.store.RecordWin("Bob")
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

// PlayerStore handles how to get and set player's score
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// GetPlayerScore return the player score
// func GetPlayerScore(name string) int {
// 	if name == "Pepper" {
// 		return 20
// 	}

// 	if name == "Floyd" {
// 		return 10
// 	}

// 	return 0
// }

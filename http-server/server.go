package httpserver

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
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// PlayerStore handles how to get and set player's score
type PlayerStore interface {
	GetPlayerScore(name string) int
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

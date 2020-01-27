package main

import (
	"log"
	"net/http"
)

type inMemoryPlayerStore struct{}

func (i *inMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *inMemoryPlayerStore) RecordWin(name string) {
	return
}

func main() {
	server := &PlayerServer{&inMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

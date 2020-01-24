package main

import (
	"log"
	"net/http"

	httpserver "github.com/lhdv/learn_go_with_tests/http-server"
)

func main() {
	server := &httpserver.PlayerServer{}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

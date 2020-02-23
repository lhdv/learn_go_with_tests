// https://quii.gitbook.io/learn-go-with-tests/build-an-application/time#just-enough-information-on-poker

package main

import (
	"log"
	"net/http"

	poker "github.com/lhdv/learn_go_with_tests/http-server"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server, _ := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

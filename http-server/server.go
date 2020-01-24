package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer server to handle GET/POST requests
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprint(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprint(w, "10")
		return
	}

	return

}

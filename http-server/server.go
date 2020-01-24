package httpserver

import (
	"fmt"
	"net/http"
)

// PlayerServer server to handle GET/POST requests
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store interface to fetch data
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server HandlerFunction
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}

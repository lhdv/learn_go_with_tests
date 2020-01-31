package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

//
// Integration tests
//
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newGetScoreRequest("Pepper"))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newRequestLeague())

		assertStatus(t, response.Code, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})

	t.Run("1000 parallel POST calls should return score 1000", func(t *testing.T) {
		calls := 1000
		score := "1000"
		// store := NewInMemoryPlayerStore()
		database, cleanDatabase := createTempFile(t, `[]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		server := NewPlayerServer(store)
		player := "Bob"

		var wg sync.WaitGroup
		wg.Add(calls)

		for i := 0; i < calls; i++ {
			go func(w *sync.WaitGroup) {
				server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
				w.Done()
			}(&wg)
		}

		wg.Wait()

		request := newGetScoreRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), score)
	})
}

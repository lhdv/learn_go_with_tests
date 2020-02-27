package poker_test

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	poker "github.com/lhdv/learn_go_with_tests/http-server"
)

//
// Integration tests
//
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()

	store, err := poker.NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	server := mustMakePlayerServer(t, store, dummyGame)

	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newGetScoreRequest("Pepper"))

		assertStatus(t, response, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(response, newRequestLeague())

		assertStatus(t, response, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := []poker.Player{
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

		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		server := mustMakePlayerServer(t, store, dummyGame)
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

		assertStatus(t, response, http.StatusOK)
		assertResponseBody(t, response.Body.String(), score)
	})
}

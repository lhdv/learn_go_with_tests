package syncing

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incremeting the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got: %d, want: %d", counter.Value(), 3)
		}
	})
}

func assertCounter(t *testing.T, got Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got: %d, want: %d", got.Value(), want)
	}
}

package selecting

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.theverge.com"
	fastURL := "http://www.google.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

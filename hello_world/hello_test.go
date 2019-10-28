package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World of Testing!!!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

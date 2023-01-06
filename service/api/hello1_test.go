package api

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world 1"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

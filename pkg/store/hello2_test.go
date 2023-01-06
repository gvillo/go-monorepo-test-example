package store

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world 2"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import "testing"

func TestHello(t *testing.T) {

	assertHelper := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	}

	t.Run("Hello function works with argument", func(t *testing.T) {
		got := Hello("Lou")
		want := "Hello, Lou"
		assertHelper(t, got, want)
	})

	t.Run("Hello function works with default argument", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertHelper(t, got, want)
	})
}

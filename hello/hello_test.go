package main

import "testing"

func TestGreeting(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say a greeting to a person", func(t *testing.T) {
		got := greeting("Fred")
		want := "Hello Fred"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to World", func(t *testing.T) {
		got := greeting("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}

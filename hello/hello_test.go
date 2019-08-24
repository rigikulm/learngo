package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting("Fred")
	want := "Hello Fred"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

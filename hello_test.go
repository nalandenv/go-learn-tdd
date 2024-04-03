package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nikhil")
	want := "Hello, Nikhil"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

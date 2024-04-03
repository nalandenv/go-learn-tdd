package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Mister", "Spanish")
		want := "Hola, Mister"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Ervin", "French")
		want := "Bonjour, Ervin"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Greeting in Marathi", func(t *testing.T) {
		got := Hello("Nikhil", "Marathi")
		want := "Namaskar, Nikhil"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() //this is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

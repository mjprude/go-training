package main

import "testing"

func assertCorrectMessage(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Michael", "English")
		want := "Hello, Michael"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Señora", "Spanish")
		want := "Hola, Señora"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mademoiselle", "French")
		want := "Bonjour, Mademoiselle"
		assertCorrectMessage(t, got, want)
	})
}

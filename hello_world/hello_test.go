package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Matt", "English")
		want := "Hello, Matt"

		assertCorrectMesage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMesage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMesage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Francois", "French")
		want := "Bonjour, Francois"

		assertCorrectMesage(t, got, want)
	})
}

func assertCorrectMesage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		expected := "Hello, Chris"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, World"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("in Franch", func(t *testing.T) {
		got := Hello("Marc", "French")
		expected := "Bonjour, Marc"

		assertCorrectMessage(t, got, expected)
	})
}

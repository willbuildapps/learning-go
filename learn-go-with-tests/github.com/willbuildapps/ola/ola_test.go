package ola

import (
	"testing"
)

func TestOla(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in spanish", func (t * testing.T) {
		t.Helper()
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t * testing.T) {
		t.Helper()
		got := Hello("William", "french")
		want := "Bonjour, William"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func (t * testing.T) {
		got := Hello("William", "")
		want := "Hello, William"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when a empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}
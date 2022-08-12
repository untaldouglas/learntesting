package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Douglas", "")
		want := "Hello, Douglas"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Maria", "Spanish")
		want := "Hola, Maria"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lola", "French")
		want := "Bonjour, Lola"
		assertCorrectMessage(t, got, want)
	})

}

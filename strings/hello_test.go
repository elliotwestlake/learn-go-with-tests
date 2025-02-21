package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello(Params{Name: "Steve"})
		want := "Hello, Steve!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should return default greeting", func(t *testing.T) {
		got := Hello(Params{})
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should return greeting in Spanish", func(t *testing.T) {
		got := Hello(Params{Name: "Elodie", Language: "Spanish"})
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should return greeting in French", func(t *testing.T) {
		got := Hello(Params{Name: "Elodie", Language: "French"})
		want := "Bonjour, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should return greeting in German", func(t *testing.T) {
		got := Hello(Params{Name: "Elodie", Language: "German"})
		want := "Hallo, Elodie!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

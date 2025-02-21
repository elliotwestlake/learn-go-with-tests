package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello(Params{Name: "Steve"})
		want := "Hello, Steve!"
		assert.Equal(t, got, want)
	})

	t.Run("should return default greeting", func(t *testing.T) {
		got := Hello(Params{})
		want := "Hello, World!"
		assert.Equal(t, got, want)
	})

	t.Run("should return greeting in Spanish", func(t *testing.T) {
		got := Hello(Params{Name: "Elodie", Language: "Spanish"})
		want := "Hola, Elodie!"
		assert.Equal(t, got, want)
	})

	t.Run("should return greeting in French", func(t *testing.T) {
		got := Hello(Params{Name: "Elodie", Language: "French"})
		want := "Bonjour, Elodie!"
		assert.Equal(t, got, want)
	})

	t.Run("should return greeting in German", func(t *testing.T) {
		got := Hello(Params{Name: "Elodie", Language: "German"})
		want := "Hallo, Elodie!"
		assert.Equal(t, got, want)
	})
}

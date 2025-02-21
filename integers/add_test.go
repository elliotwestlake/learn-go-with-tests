package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("should return sum of two integers", func(t *testing.T) {
		got := Add(1, 2)
		want := 3

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

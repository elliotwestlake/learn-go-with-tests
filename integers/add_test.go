package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("should return sum of two integers", func(t *testing.T) {
		got := Add(1, 2)
		want := 3

		assert.Equal(t, got, want)
	})
}

func ExampleAdd() {
	fmt.Println(Add(1, 2))
	// Output: 3
}

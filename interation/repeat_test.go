package interation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

	assert.Equal(t, got, want)
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 10))
	// Output: aaaaaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

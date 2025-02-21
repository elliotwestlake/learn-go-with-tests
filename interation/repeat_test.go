package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"
	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
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

package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 7)
	expected := "aaaaaaa"

	if got != expected {
		t.Errorf("expected '%s' but got '%s'", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("x", 8))
	fmt.Println(Repeat("z", 1))
	fmt.Println(Repeat("t", 0))

	// Output: xxxxxxxx
	// z
	// 
}

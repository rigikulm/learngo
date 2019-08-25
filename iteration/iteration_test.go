package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	count := 8
	want := "aaaaaaaa"
	got := Repeat("a", count)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func ExampleRepeat() {
	s := Repeat("x", 7)
	fmt.Println(s)
	// Output: xxxxxxx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("c", 10)
	}
}

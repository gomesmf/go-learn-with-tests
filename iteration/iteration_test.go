package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestJoin(t *testing.T) {
	s := []string{"my", "name", "is", "Matheus"}
	got := Join(s, "-")
	expected := "my-name-is-Matheus"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	// Output: aaaaaa
}

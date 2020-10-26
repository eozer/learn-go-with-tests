package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	letters := Repeat("b", 3)
	fmt.Println(letters)
	// Output: bbb
}

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 5)
	expected := "aaaaa"
	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a", 10)
	}
}

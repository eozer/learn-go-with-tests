package integers

import (
	"fmt"
	"testing"
)

// This is an example to be added to go doc. When we do not put Output directive it will not run,
// even if it is compiled. Try it with $ go test -v
func ExampleAdd() {
	sum := Add(100, 1000)
	fmt.Println(sum)
	// Output: 1100
}

func TestAdder(t *testing.T) {
	actual := Add(2, 5)
	expected := 7
	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers, fixed (nope) size", func(t *testing.T) {
		// NOTE: fixed size array != slice, so we must give exactly the needed type
		array := []int{1, 2, 3, 4, 5}
		actual := Sum(array)
		expected := 15
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("sum of numbers,  any size", func(t *testing.T) {
		array := []int{1, 2, 3}
		actual := Sum(array)
		expected := 6
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

}

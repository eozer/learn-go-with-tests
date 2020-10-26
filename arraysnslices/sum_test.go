package sum

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("given 1 slice", func(t *testing.T) {
		actual := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
		expected := []int{15, 6}
		// NOTE: we cannot compare slices like actual != expected, we need to use:
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})

	t.Run("given 2 slices", func(t *testing.T) {
		actual := SumAll([]int{1, 2, 4})
		expected := []int{7}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})
}

func TestSumTails(t *testing.T) {
	t.Run("sum tails", func(t *testing.T) {
		actual := SumAllTails([]int{4, 5}, []int{2, 3})
		expected := []int{5, 3}
		// NOTE: we cannot compare slices like actual != expected, we need to use:
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})
	t.Run("2nd test case sum tails", func(t *testing.T) {
		actual := SumAllTails([]int{4, 5, 6}, []int{2, 3, 1})
		expected := []int{11, 4}
		// NOTE: we cannot compare slices like actual != expected, we need to use:
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})
	t.Run("empty slice, sum tails", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{2, 3, 1})
		expected := []int{0, 4}
		// NOTE: we cannot compare slices like actual != expected, we need to use:
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	})
}

package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("computes Perimeter", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := Perimeter(r)
		want := 40.0
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("computes Area of a rectangle", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := Area(r)
		want := 100.0
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("computes Area of a circle", func(t *testing.T) {
		c := Circle{10}
		got := Area(c)
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}

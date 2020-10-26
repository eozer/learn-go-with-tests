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

	checkArea := func(t *testing.T, s Shape, want float64) {
		got := s.Area()
		if got != want {
			t.Errorf("%#v got %g, want %g", s, got, want)
		}
	}

	t.Run("computes Area of a rectangle", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		checkArea(t, r, 100.0)
	})

	t.Run("computes Area of a circle", func(t *testing.T) {
		c := Circle{10}
		checkArea(t, c, 314.1592653589793)
	})

	t.Run("table driven tests: rectangle, circle, triangle...", func(t *testing.T) {
		tests := []struct {
			shape   Shape
			hasArea float64
		}{
			{shape: Rectangle{10.0, 10.0}, hasArea: 100.0},
			{shape: Circle{10.0}, hasArea: 314.1592653589793},
			{shape: Triangle{12, 6}, hasArea: 36.0},
		}
		for _, test := range tests {
			checkArea(t, test.shape, test.hasArea)
		}
	})

}

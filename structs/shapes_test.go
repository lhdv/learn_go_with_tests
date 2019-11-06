package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		want := 72.0

		checkArea(t, rect, want)
	})

	t.Run("circles", func(t *testing.T) {
		circ := Circle{10}
		want := 314.1592653589793

		checkArea(t, circ, want)
	})
}

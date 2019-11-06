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

	checkTest := func(t *testing.T, got, want float64) {
		t.Helper()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		got := rect.Area()
		want := 72.0

		checkTest(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circ := Circle{10}
		got := circ.Area()
		want := 314.1592653589793

		checkTest(t, got, want)
	})
}

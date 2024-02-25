package structMethod

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		checkPerimeter(t, Rectangle{10.0, 10.0}, 40.0)
	})

	t.Run("circles", func(t *testing.T) {
		checkPerimeter(t, Circle{10.0}, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		checkArea(t, Rectangle{12.0, 6.0}, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		checkArea(t, Circle{10.0}, 314.1592653589793)
	})
}

func checkArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func checkPerimeter(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

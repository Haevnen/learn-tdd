package structMethod

import "testing"

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{10.0, 10.0}, 40.0},
		{"circle", Circle{10.0}, 62.83185307179586},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.shape, tt.want)
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"circle", Circle{10.0}, 314.1592653589793},
		{"triangle", Triangle{3, 4}, 6.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
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

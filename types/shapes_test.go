package types

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.00, 10.00}
	got := Perimeter(rectangle)
	want := 40.00
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	t.Run("Area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.00, 10.00}
		got := rectangle.Area()
		want := 100.00
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("Area of circle", func(t *testing.T) {
		circle := Circle{10.00}
		got := circle.Area()
		want := 314.00
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}

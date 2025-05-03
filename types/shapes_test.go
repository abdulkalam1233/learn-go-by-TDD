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
	rectangle := Rectangle{10.00, 10.00}
	got := Area(rectangle)
	want := 100.00
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

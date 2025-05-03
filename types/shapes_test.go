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

	//checkArea := func(t *testing.T, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %.2f, want %.2f", got, want)
	//	}
	//}
	//
	//t.Run("Area of rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{10.00, 10.00}
	//	want := 100.00
	//	checkArea(t, rectangle, want)
	//})
	//
	//t.Run("Area of circle", func(t *testing.T) {
	//	circle := Circle{10.00}
	//	want := 314.1592653589793
	//	checkArea(t, circle, want)
	//})

	areaTests := []struct {
		Shape Shape
		want  float64
	}{
		{Rectangle{10.00, 10.00}, 100.00},
		{Circle{10.00}, 314.1592653589793},
		{Triangle{10.00, 10.00}, 50.00},
	}
	for _, test := range areaTests {
		got := test.Shape.Area()
		if got != test.want {
			t.Errorf("got %.2f, want %.2f", got, test.want)
		}
	}
}

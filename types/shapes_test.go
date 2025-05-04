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

	//checkArea := func(t *testing.T, shape Shape, hasArea float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != hasArea {
	//		t.Errorf("got %.2f, hasArea %.2f", got, hasArea)
	//	}
	//}
	//
	//t.Run("Area of rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{10.00, 10.00}
	//	hasArea := 100.00
	//	checkArea(t, rectangle, hasArea)
	//})
	//
	//t.Run("Area of circle", func(t *testing.T) {
	//	circle := Circle{10.00}
	//	hasArea := 314.1592653589793
	//	checkArea(t, circle, hasArea)
	//})

	// Table tests
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{10.00, 10.00}, hasArea: 100.00},
		{name: "Circle", shape: Circle{10.00}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10.00, 10.00}, hasArea: 50.00},
	}
	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.hasArea {
			// clear assertions
			t.Errorf("%#v got %g want %g", test.shape, got, test.hasArea)
		}
	}
}

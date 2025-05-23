package types

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

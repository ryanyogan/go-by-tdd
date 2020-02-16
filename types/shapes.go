package types

import "math"

// Shape defines the methods available on Rectangles, Circles, and
// Triangles.
type Shape interface {
	Area() float64
}

// Rectangle holds the height and width fields of a Rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area takes a height and a width, returns the area of the values.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle holds the radius field of a Circle.
type Circle struct {
	Radius float64
}

// Area takes a Radius from a Circle, returns the area of the values.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle holds the base and height fields of a triangle.
type Triangle struct {
	Base   float64
	Height float64
}

// Area reutrns the area of a Triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

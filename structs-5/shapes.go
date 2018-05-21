package shapes

import "math"

// Rectangle is a type of shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a type of shape
type Circle struct {
	Radius float64
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

// Perimeter returns the perimeter of a rectangle given a width and height
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

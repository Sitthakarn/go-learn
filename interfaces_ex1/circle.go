package main

import "math"

// Circle Structure
type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) diameter() float64 {
	return 2 * c.Radius
}

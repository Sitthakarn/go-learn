package main

// Square Structure
type Square struct {
	sideLenght float64
}

func (s Square) area() float64 {
	return s.sideLenght * s.sideLenght
}

func (s Square) perimeter() float64 {
	return s.sideLenght * 4.0
}

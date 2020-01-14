package main

// MyDrawing structure
type MyDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

// Export the Area() function of MyDrawing
func (drawing MyDrawing) area() float64 {
	totalArea := 0.0
	for _, s := range drawing.shapes {
		totalArea += s.area()
	}
	return totalArea
}

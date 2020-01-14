package main

import "fmt"

// Print all shape properties in range
func printShape(shapes ...Shape) {
	for _, shape := range shapes {
		fmt.Printf("Shape Type = %T, Shape Value = %v\n", shape, shape)
		fmt.Printf("Area = %f\n", shape.area())
		fmt.Printf("Perimeter = %f\n", shape.perimeter())
	}

}

// Calculate the all shape area
func calculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}

func main() {
	var c Shape = Circle{5.0}
	//var s Shape = Square{5.0}

	printShape(c, Square{5.0})
	fmt.Printf("Total area = %f\n", calculateTotalArea(c, Square{5}))
	fmt.Printf("\n")

	drawing := MyDrawing{
		shapes: []Shape{
			Circle{Radius: 5},
			Square{sideLenght: 5},
		},
		bgColor: "White",
		fgColor: "red",
	}

	fmt.Println("Drawing", drawing)
	fmt.Println("Total area = ", drawing.area())

}

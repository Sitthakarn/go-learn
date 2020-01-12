package main

import "fmt"

func main() {
	colors := make(map[int]string)

	colors[10] = "#ff0000"
	colors[11] = "#ed4393"

	fmt.Println(colors)
	delete(colors, 11)
	fmt.Println(colors)

	colors2 := map[string]string{
		"red":   "#ff0000",
		"green": "#4df4267",
		"white": "#ffffff",
	}
	printMap(colors2)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

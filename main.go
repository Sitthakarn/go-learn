package main

import "fmt"

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	fmt.Println(cards.toString())

	cards.saveToFile("mycards.crd")
}

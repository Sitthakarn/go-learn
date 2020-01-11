package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// fmt.Println(cards.toString())

	// cards.saveToFile("mycards2.crd")
	cards := newDeckFromFile("mycards2.crd")
	cards.print()
	cards.shuffle()
	cards.print()
	cards.saveToFile("mycards3.crd")
}

package main

func main() {
	cards := newDeck()
	cards.saveToFile("mycards2.crd")

	// hand, remainingCards := deal(cards, 5)

	cards = newDeckFromFile("mycards2.crd")
	cards.print()
	cards.shuffle()
	cards.print()
	cards.saveToFile("mycards3.crd")
}

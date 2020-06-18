package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	hand, deck := deal(cards, 5)
	hand.print()
	deck.print()
}

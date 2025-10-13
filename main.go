package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 8)

	hand.print()
	remainingCards.print()
}



package main

func main() {
	cards := deck{"Seven of Hearts", newCard()}
	cards = append(cards, "Four of Spaces")

	cards.print()
}

func newCard() string {
	return "Nine of Clubs"
}

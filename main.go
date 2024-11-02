package main

func main() {
	cards := deck{"Seven of Diamonds", newCard()}
	cards = append(cards, "Six of Spaces")

	cards.print()
}

func newCard() string {
	return "Eight of Clubs"
}

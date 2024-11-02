package main

import "fmt"

func main() {
	cards := []string{"Seven of Diamonds", newCard()}
	cards = append(cards, "Six of Spaces")

	fmt.Println(cards)
}

func newCard() string {
	return "Eight of Clubs"
}

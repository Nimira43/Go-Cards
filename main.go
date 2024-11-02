package main

import "fmt"

func main() {
	cards := []string{"Seven of Diamonds", newCard()}
	cards = append(cards, "Six of Spaces")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Eight of Clubs"
}

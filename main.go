package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spaces")


}

func newCard() string {
	return "Five of Diamonds"
}
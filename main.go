package main

import "fmt"

func main() {
  card := newCard()
	fmt.Println(card)
}

func newCard() {
	return "Five of Diamonds"
}
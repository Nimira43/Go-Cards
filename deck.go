package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
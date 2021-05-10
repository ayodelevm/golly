package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// (d deck) - receiver helps attach a method to a variable of the same type...
// d is the value of the variable method is called on accessble within the function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

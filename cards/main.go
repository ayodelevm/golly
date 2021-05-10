package main

// import "fmt"

func main() {
	// var card string = "This is a card var"
	card := newDeck()

	hand, remainingDeck := deal(card, 5)
	card.print()
	hand.print()
	remainingDeck.print()
}

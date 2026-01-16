package main

import "fmt"

func main() {
	// Create a new variable in standard way with type declaration
	// var card string = "Ace of Spades"
	// Create a new variable in short way (Make the Go compiler determine the type)
	// := is only used upon the first declaration of a variable
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()
	// cards := deck{"Ace of Diamonds", newCard()}
	// // Append returns a new slice and does not modify the original slice
	// cards = append(cards, "Six of Spades")

	// fmt.Println(card)
	// fmt.Println(cards)

	// Standard way
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards := newDeck()
	cards.print()

	hand, reaminingDeck := deal(cards, 5)

	hand.print()
	reaminingDeck.print()

	fmt.Println(reaminingDeck.toString())

	reaminingDeck.shuffle()

	reaminingDeck.saveToFile("my_cards")

	newDeck := newDeckFromFile("my_cards")
	newDeck.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }

package main

import "fmt"

func main() {
	deckOfCard := getNewDeck()
	fmt.Print("Original Deck ---->")
	deckOfCard.printDeck()
	fmt.Println()
	var remainingDeck deck
	deckOfCard, remainingDeck = deal(deckOfCard, 3)
	fmt.Println()
	fmt.Print("Remaining Deck ---->")
	remainingDeck.printDeck()
	fmt.Println()
	fmt.Print("Dealt Deck ---->")
	deckOfCard.printDeck()
	fmt.Println()

}

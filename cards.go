package main

import "fmt"

func main() {

	deckOfCard := deck{}
	deckOfCard = append(deckOfCard, "king")
	deckOfCard = append(deckOfCard, "Queen")
	deckOfCard = append(deckOfCard, "jack")

	deckOfCard.printDeck()
	fmt.Print("hello Kushagra")

}

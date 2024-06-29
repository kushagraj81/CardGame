package main

import "fmt"

func main() {

	deckOfCard := deck{}
	deckOfCard = append(deckOfCard, "king")
	deckOfCard = append(deckOfCard, "Queen")
	deckOfCard = append(deckOfCard, "jack")

	printDeck(deckOfCard)

	fmt.Print("hello Kushagra")

}

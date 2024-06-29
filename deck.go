package main

import "fmt"

type deck []string

// If we declare function like this then it adds the function to the new datatype we created and we can call deckOfObject.printDeck()
func (d deck) printDeck() {
	// for i, card := range d {
	// 	fmt.Println(i, card)
	// }

	for i := 0; i < len(d); i++ {
		fmt.Print(d[i] + " ## ")
		if i == len(d)-1 {
			fmt.Println()
		}
	}
}

// Above if we decalred function like  func printDeck(d deck) {
// Then it would have added function to class and we could not have used it on object of deck type deckOfObject.printDeck()
// then we wouuld have to use it like printDeck( deckOfObject)

func getNewDeck() deck {
	cards := deck{}

	suits := deck{"Spade", "Heart"}
	values := deck{"Ace", "King"}

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			cards = append(cards, values[j]+" Of "+suits[i])
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

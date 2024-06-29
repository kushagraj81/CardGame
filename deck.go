package main

import "fmt"

type deck []string

func printDeck(d deck) {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

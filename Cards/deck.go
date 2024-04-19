// code that describes what a deck is and how it works

package main

import "fmt"

// create  a new type of 'deck', (we have created our own type and named it 'deck')
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

// (d deck) is a receiver here, any variable of type deck gets access to print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

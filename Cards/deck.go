// code that describes what a deck is and how it works

package main

import (
	"fmt"
	"os"
	"strings"
)

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


// return two separate values in a single function
// here the function is returning two values each of type deck
// syntax of a function is : func keyword then name of the function (arguments) return type o fthe function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
  return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// err is a value of type error, if anything goes wrong there will be some error value, if not everything goe sright then err will have value "nil"

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file: ", err)
    os.Exit(1)
	}
  s := strings.Split(string(bs), ",")
  return deck(s)
}

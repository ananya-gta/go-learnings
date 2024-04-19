// Note: files in the same package do not have to be imported into each other

package main

import "fmt"

/*
func main() {
	// Go is a static type of programming language, you need to specify data types to variables just like java
	// var card string = "Ace of Hearts" or use card := "Ace of Hearts"
	// use colon only for very first initialization
	card := newCard()
	fmt.Println("Card: ", card)
}

func newCard() string{
	return "Five of Diamonds"
} */

// Array: Fixed Length list of things of same data type
// Slice: An array that can grow or shrink, length is not fixed

func main() {
	cards := []string{"Ace of Hearts", newCard()}
	cards = append(cards, "Five of Diamonds") //slice

	// iterating over elements
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Six of Spades"

}

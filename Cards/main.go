// Note: files in the same package do not have to be imported into each other

package main

import "fmt"

func main() {
	// Go is a static type of programming language, you nee dto specify data types to variables just like java
	// var card string = "Ace of Hearts" or use card := "Ace of Hearts"
	// use colon only for very first initialization
	card := newCard() 
	fmt.Println("Card: ", card)
}

func newCard() string{
	return "Five of Diamonds"
}

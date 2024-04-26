package main

import (
	"fmt"
)

type bot interface {
  getGreetings() string
}

// creating these structs so that i can create some functions that will use these structs as recievers
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
  printGreeting(sb)
}


func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

// since we are not using value of the receiver we can remove the eb and sb
func (eb englishBot) getGreetings() string {
	// VERY custom logic for generating english greeting
	return "Hello"
}

func (sb spanishBot) getGreetings() string {
	return "Hola"
}

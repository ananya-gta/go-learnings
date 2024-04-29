package main

import (
	"fmt"
)

type Printer interface {
	Print()
}

type StringPrinter struct{}

func (s StringPrinter) Print() {
	fmt.Println("StringPrinter")
}

type IntPrinter struct{}

func (i IntPrinter) Print() {
	fmt.Println("IntPrinter")
}

func main() {
	printers := []Printer{StringPrinter{}, IntPrinter{}}
	for _, printer := range printers {
		printer.Print()
	}
}

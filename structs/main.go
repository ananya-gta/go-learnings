package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// also yo ucan embed a struct in this way
// type person struct {
//   firstName string
//   lastNmae string
//    contactInfo
// }

func main() {
	// declaring struct
	// 1st way
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	//  fmt.Println(alex)

	//2nd way
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	//3rd way
	jim := person{
		firstName: "Jim",
		lastName:  "Andrew",
		contact:   contactInfo{email: "jim.andrew@gmail.com", zipcode: 208011},
	}

  jimPointer := &jim
  jimPointer.updateFirstName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// use of go pointers, go is a pass by value language
func (pointerToPerson *person) updateFirstName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}

// address --> value := *address
// value --> address := &value

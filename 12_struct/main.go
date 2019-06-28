package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " +
		strconv.Itoa(p.age)
}

func main() {
	// init person using struct
	person1 := Person{
		firstName: "Bode",
		lastName:  "Fals",
		city:      "Minneapolis",
		gender:    "M",
		age:       40}

	// Alternative
	// person1 := Person{"Bode", "Fals", "Minneapolis", "M", 40}
	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	fmt.Println(person1.greet())
}

package main

import ( // struct
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

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}

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
	person2 := Person{"Sheri", "Taiwo", "Minneapolis", "F", 30}
	fmt.Println(person2)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)
	person1.hasBirthday()
	person2.getMarried("Fals")
	fmt.Println(person2.greet())
}

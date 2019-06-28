package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string) // defines a map

	// Assign kv
	emails["Bode"] = "bode40@hotmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bode"])
	fmt.Println(len(emails))

	// Declare map and add kv
	nemails := map[string]string{"Bob": "bob@gmail.com", "Susie": "susie@gmail.com"}
	fmt.Println(nemails)

	// Delete from map
	delete(emails, "Bode")
	fmt.Println(emails)
}

package main

import "fmt"

func main() { // using variables and types

	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Bode"
	var age = 37
	const isCool = true
	var sizes float32 = 2.3

	// Shorthand
	// name := "Bode" // shorthand declaration
	// email := "bode40@hotmail.com"

	name, email := "Bode", "bode40@hotmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", isCool) // %T prints out the type of the argumen
	fmt.Printf("%T\n", sizes)

}

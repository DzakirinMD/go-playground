package main

import "fmt"

// accepting in a pointer
func updateNamePointer(x *string) {
	*x = "wedge"
}

// &
func PointersLearn() {
	name := "tifa"

	// This will show the memory location
	fmt.Println("memory address of name is : ", &name)

	// this will store the pointer(memory address) of name into variable m
	m := &name
	fmt.Println("memory address: ", m)
	// Dereference. it will check back whats the value at the memory address
	fmt.Println("value at the memory address: ", m, " is : ", *m)

	updateNamePointer(m)

	fmt.Println(name)
}

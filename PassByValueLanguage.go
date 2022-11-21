package main

import "fmt"

/*
Go makes "copies" of values when passed into functions
*/

func updateName(x string) {
	x = "wedge"
}

func updateNameReal(n string) string {
	n = "wedge"
	return n
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func passByValue() {

	/*
		Group A (Non-Pointer Values) types -> strings, ints, bools, float, arrays, struct
		This group doesn't actually change the value of the original. It change the value of the copy in the memory
	*/
	name := "tifa"

	/*
		Group B (Pointer Wrapper Values) types -> Slices, Maps, Functions
		This group will actually change the original value. because it will refer the value of the original pointer
	*/
	menu := map[string]float64{
		"soup": 4.99,
		"pie":  7.99,
	}

	// It actually update the value of the copy in the memory. The original value is not updated
	updateName(name)

	// This is why we still see tifa in this Println
	fmt.Println(name)

	name = updateNameReal(name)

	fmt.Println(name)

	updateMenu(menu)

	fmt.Println(menu)
}

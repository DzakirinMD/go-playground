package main

import "fmt"

/*
struct is a blueprint of data
basically like entity in JAVA
it can have method
*/
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"soup": 4.99, "pie": 7.99},
		tip:   0,
	}

	return b
}

// Receiver functions. Now this method only can be call by bill object
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	// the %-25v meaning it is 25 character long
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

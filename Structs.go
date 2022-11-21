package main

import (
	"fmt"
	"os"
)

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
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	// the %-25v meaning it is 25 character long
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	// tip
	fs += fmt.Sprintf("%-25v ...$%0.2f", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "\ntotal:", total+b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	// turn to byte
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved to fill")
}

// use Userinput.go and Struct.go
// go program to create a bill system
package bills

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Create bill
func NewBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"soup": 4.99, "pie": 7.99},
		tip:   0,
	}

	return b
}

// add an item to the bill
func (b *bill) AddItem(name string, price float64) {
	b.items[name] = price
}

// update tip
func (b *bill) UpdateTip(tip float64) {
	b.tip = tip
}

// Save bill
func (b *bill) Save() {
	// turn to byte
	data := []byte(b.Format())

	err := os.WriteFile("mini-project/bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved to fill")
}

// Receiver functions. Now this method only can be call by bill object
func (b *bill) Format() string {
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

package bills

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read from terminal
func CreateBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := GetInput("Create a new bill name : ", reader)

	b := NewBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

// helper function
func GetInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, error := reader.ReadString('\n')

	return strings.TrimSpace(input), error
}

func PromtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := GetInput("Choose options (a - add item, s - Save bill, t - add tip) : ", reader)

	switch option {
	case "a":
		name, _ := GetInput("Item name: ", reader)
		price, _ := GetInput("Item price: ", reader)

		// parsing string to float64
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			PromtOptions(b)
		}

		b.AddItem(name, p)
		fmt.Println("item added - ", name, price)
		PromtOptions(b)
	case "s":
		b.Save()
		fmt.Println("you Save the file - ", b.name)
	case "t":
		tip, _ := GetInput("Enter tip amount ($): ", reader)
		// parsing string to float64
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			PromtOptions(b)
		}
		b.UpdateTip(t)
		fmt.Println("tip added - ", tip)
		PromtOptions(b)

	default:
		fmt.Println("Not a valid option...")
		// this will call the function again to re-choose
		PromtOptions(b)
	}
}

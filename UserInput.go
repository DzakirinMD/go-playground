package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read from terminal
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name : ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

// helper function
func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, error := reader.ReadString('\n')

	return strings.TrimSpace(input), error
}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose options (a - add item, s - save bill, t - add tip) : ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		// parsing string to float64
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promtOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("item added - ", name, price)
		promtOptions(b)
	case "s":
		b.save()
		fmt.Println("you save the file - ", b.name)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		// parsing string to float64
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promtOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promtOptions(b)

	default:
		fmt.Println("Not a valid option...")
		// this will call the function again to re-choose
		promtOptions(b)
	}
}

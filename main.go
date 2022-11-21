package main

import (
	"fmt"
	"go_playground/bills"
)

// this variable become packaged scope
var score = 99.5

func main() {

	// when declared here the PackageScope.go cannot access here because this score belong to main()
	// var score = 99.5

	sayHello("mario")
	for _, value := range points {
		fmt.Println(value)
	}

	showScore()

	Maps()

	passByValue()

	PointersLearn()

	// Structs (Entity)
	mybill := newBill("marios's bill")

	mybill.updateTip(10)
	mybill.addItem("coffee", 4.99)
	mybill.addItem("pudding", 10)

	fmt.Println(mybill)

	billObject := bill{
		name:  "name",
		items: map[string]float64{},
		tip:   0,
	}

	fmt.Println(billObject)

	fmt.Println(mybill.format())

	// ########################  Bill part - start ########################
	billing := bills.CreateBill()

	bills.PromtOptions(billing)

	fmt.Println(billing)

	// ######################## Bill part end - End ########################

	// ######################## Interface - Start ########################
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, value := range shapes {
		printShapeInfo(value)
		fmt.Println("---")
	}
	// ######################## Interface - End ########################
}

package main

import "fmt"

// this variable become packaged scope
var score = 99.5

func main() {

	// when declared here the PackageScope.go cannot access here because this score belong to main()
	// var score = 99.5

	// sayHello("mario")
	// for _, value := range points {
	// 	fmt.Println(value)
	// }

	// showScore()

	// Maps()

	// passByValue()

	// PointersLearn()

	// Structs (Entity)
	mybill := newBill("marios's bill")

	fmt.Println(mybill)

	billObject := bill{
		name:  "name",
		items: map[string]float64{},
		tip:   0,
	}

	fmt.Println(billObject)

	fmt.Println(mybill.format())
}

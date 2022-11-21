package main

import (
	"fmt"
	"math"
)

// nameFour := "test"

func Variables() {

	// variables
	var nameOne string = "mario"
	// it will auto infer
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println("nameOne = " + nameOne + " nameTwo = " + nameTwo + " nameThree = " + nameThree)
	fmt.Println("nameOne = ", nameOne, " nameTwo = ", nameTwo, " nameThree = ", nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println("nameOne = ", nameOne, " nameTwo = ", nameTwo, " nameThree = ", nameThree)

	// shorthand version for initializing variable. HOWEVER DECLARE LIKE THIS you cant use it outside of function
	nameFour := "yoshi"

	fmt.Println("nameOne = ", nameOne, " nameTwo = ", nameTwo, " nameThree = ", nameThree, " nameFour = ", nameFour)

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// ints
	var ageOne int = 20
	var ageTwo int = 30
	ageThree := 40

	fmt.Println("ageOne = ", ageOne, " ageTwo = ", ageTwo, " ageThree = ", ageThree)

	// bits & memory
	var numOne int8 = 127
	var numTwo int16 = -128
	// uint is only positive number, cannot be negative number
	var numThree uint8 = 255

	// float. most time use float64
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8182381294313124986786.66

	fmt.Println("numOne = ", numOne,
		" numTwo = ", numTwo,
		" numThree = ", numThree,
		" scoreOne = ", scoreOne,
		" scoreTwo = ", scoreTwo)
}

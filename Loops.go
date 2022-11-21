package main

import "fmt"

func Loops() {
	x := 0

	// same like while x < 5
	for x < 5 {
		fmt.Println("value of x is : ", x)
		x++
	}

	for i := 0; i <= 5; i++ {
		fmt.Println("value of i is : ", i)
	}

	names := []string{"mario", "luigi", "yoshi", "bowser"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// cycle thru slices/array. getting the index and value at that index
	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}
	// only if dont want to use then put underscode. below will only get the value
	// alter value in here wont alter the original value
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}
}

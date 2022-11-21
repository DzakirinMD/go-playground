package main

import "fmt"

func BooleanAndCondition() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	// if, else if, else
	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	// nested if, continue, break
	names := []string{"mario", "luigi", "yoshi", "yoshi", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos ", index)
			//  continue statement forces the next iteration of the loop to take place, skipping any code in between.
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos ", index)
			// once it enter break. the loop stop at this index
			break
		}
		// at index = 1 this will not be print because of continue
		fmt.Printf("The value at index %v is %v \n", index, value)
	}
}

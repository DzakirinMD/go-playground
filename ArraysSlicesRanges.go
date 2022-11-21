package main

import "fmt"

func ArraySlicesRanges() {

	// the array length is fixed when created
	// 2 way of declaring array
	var age [3]int = [3]int{22, 56, 31}
	var ages = [3]int{22, 56, 31}

	names := [4]string{"jack", "boss", "hero", "peach"}

	fmt.Println(ages, len(age))
	fmt.Println(names, len(names))

	/*
		SLICES
		use array under the hood. but more flexible
	*/

	var scores = []int{100, 50, 60}
	scores[1] = 222
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	/*
		SLICES RANGE
		inclusive of the first number, but not the 2nd number
	*/

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}

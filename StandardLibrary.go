package main

import (
	"fmt"
	"sort"
	"strings"
)

func StandardLibrary() {

	/*
		Documentation : https://pkg.go.dev/std
	*/

	greeting := "Hi there !"

	fmt.Println(strings.Contains(greeting, "Hi"))
	fmt.Println(strings.ReplaceAll(greeting, "Hi", "hello"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "re"))
	// this will return the slice of greeting
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged when altered by string library
	fmt.Println("original string value = ", greeting)

	// sort int
	ages := []int{45, 66, 95, 56, 23, 234, 1}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 66)
	fmt.Println(index)

	// sort string
	names := []string{"abu", "kj", "zaila"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "zaila"))

}

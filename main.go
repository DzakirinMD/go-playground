package main

import (
	"fmt"
	"strings"
)

// This function will accept string name and return 2 strings
func getInitials(names string) (string, string) {
	s := strings.ToUpper(names)

	splitNames := strings.Split(s, " ")

	var initials []string

	for _, values := range splitNames {
		initials = append(initials, values[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	// return the 2 initial as underscode if the name is only 1
	return initials[0], "_"
}

func main() {
	firstName, secondName := getInitials("albert pisa")
	fmt.Println(firstName, secondName)

	firstName1, secondName1 := getInitials("tifa lockhart")
	fmt.Println(firstName1, secondName1)
}

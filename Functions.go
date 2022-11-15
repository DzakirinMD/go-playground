package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good Morning %v \n", name)
}

func sayByee(name string) {
	fmt.Printf("Good Bye %v \n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, value := range names {
		f(value)
	}
}

// if the method have return type. need to specify after ()
func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func Functions() {
	sayGreeting("Brenda")
	sayByee("Brenda")

	cycleNames([]string{"mario", "luigi", "yoshi", "bowser"}, sayGreeting)
	cycleNames([]string{"mario", "luigi", "yoshi", "bowser"}, sayByee)

	area1 := circleArea(10.5)
	area2 := circleArea(15)

	fmt.Printf("circle 1 is %0.3f and circle 2 is % 0.3f", area1, area2)
}

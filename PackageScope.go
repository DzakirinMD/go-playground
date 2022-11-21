package main

import "fmt"

/*
All of this variable, functions will be automatically available in the other package "main" to be used
*/
var points = []int{20, 90, 145, 66, 80}

func sayHello(n string) {
	fmt.Println("hello ", n)
}

func showScore() {
	fmt.Println("you cored this many points:", score)
}

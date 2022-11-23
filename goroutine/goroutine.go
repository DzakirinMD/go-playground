package main

import (
	"fmt"
	"time"
)

/*
A goroutine is a function that executes independently from the rest of the program.
The basic element of a goroutine is a function. Every function can become goroutines.
Launching a goroutine is as simple as launching a function except that you just add the “go” word just before the function call.

In this program we launch 2 goroutines sequentially.
*/

func main() {
	//panic("show me the goroutine")

	fmt.Println("launch 1st goroutine")
	go printNumber()
	fmt.Println("launch 2nd goroutine")
	go printNumber()
	//time.Sleep(1 * time.Minute)

	// ######################### SELECT ####################### //
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	// ######################### SELECT ####################### //
}

func printNumber() {
	i := 0

	// this is an infinite loop that pause for 1 second
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i)
	}

}

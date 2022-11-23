package main

import "fmt"

// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

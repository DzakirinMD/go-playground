package main

import (
	"fmt"
	"log"
	"time"
)

/*
Goroutines can communicate with each others through channels.
A channel can be seen as a pipeline of data between two goroutines.
This pipeline can only support a specific type.

1. A channel that can can be used to SEND values of type T is denoted : chan<- T

2. A channel that can can be used to RECEIVE values of type T is denoted : <-chan T

3. A channel that can can be used to SEND and RECEIVE values of type T is denoted : chan T
*/

func main() {

	// Unbuffered channel, when do not specify the buffer size. or set to 0
	//ch1 := make(chan float64)
	//ch2 := make(chan float64, 0)

	// Buffered channel, when u specify the buffer size
	ch3 := make(chan float64, 16)

	/*
		SEND data into channel
		You can send on a channel if it’s open. If you send on a closed channel your program will panic !
		If you send on a nil channel it blocks your program forever
	*/
	ch3 <- 4
	ch3 <- 56

	// RECEIVED from channel
	received := <-ch3

	// A sender can close a channel to indicate that no more values will be sent.
	close(ch3)

	// When a channel has been closed. it will cause panic exception
	// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	go testClosingChannel(ch3)

	// This syntax is used in order to be sure that our channel is not closed or empty
	// return true if the channel is closed
	received, ok := <-ch3
	if ok {
		log.Println("the channel is empty or closed")
	}

	fmt.Println("data received from channel is", received, ok)

	// Unbuffered channel (capacity = 0)
	// When you send data into an unbuffered channel your current goroutine will be blocked until
	// the data is received by another goroutine.
	ch4 := make(chan int)
	go dummy(ch4)
	log.Println("waiting for reception...")
	ch4 <- 45
	log.Println("received")

	// Buffered channel (capacity > 0)
	// When you send data into an buffered channel your current goroutine will be blocked until
	// the data is copied to the buffer.
	ch5 := make(chan int, 1)
	go dummy(ch5)
	log.Println("waiting for reception...")
	ch5 <- 45
	log.Println("received")

	// Unbuffered channels are used to synchronize two goroutines
	syncCh := make(chan bool)
	// launch a second goroutine
	go func() {
		longTask2()
		// finished
		syncCh <- true
	}()
	longTask()
	// blocks until the second goroutine has finished
	<-syncCh
}

// goroutine
func dummy(c chan int) {
	time.Sleep(3 * time.Second)
	<-c
}

func longTask2() {
	time.Sleep(1 * time.Second)
}

func longTask() {
	time.Sleep(3 * time.Second)
}

func testClosingChannel(channel chan float64) {

	_, ok := <-channel
	if ok {
		log.Println("the channel is empty or closed")
	} else {
		channel <- 100
	}
}

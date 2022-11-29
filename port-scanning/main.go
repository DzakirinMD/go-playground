package main

import (
	"fmt"
	"go_playground/port-scanning/port"
	"time"
)

/*
Port scanning is the act of iterating over every port on a machine and checking to see which ones are Open, Closed, or Filtered.

In total there are just over 130,000 ports on a typical machine, 65535 of which are TCP and the other 65535 which are UDP ports.
Each of these ports could effectively be a way in to your system if left open and port scanning allows security engineers see if there are any potential
ways to gain access to your system from un-patched software.
*/

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Println("Port Scanning")
	// uncomment this to run in serial
	results := port.InitialScan("localhost")
	fmt.Println(results)

	// uncomment this to run a wide scan
	widescanresults := port.WideScan("localhost")
	fmt.Println(widescanresults)

	end := time.Now().UnixNano() / int64(time.Millisecond)
	diff := (end - start)
	fmt.Println("Program ended in : ", diff, "ms")
	// program took 6797 ms when run in serial
	// program took 1629 ms when run in parallel
}

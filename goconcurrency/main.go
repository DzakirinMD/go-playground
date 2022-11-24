package main

import (
	"encoding/json"
	"go_playground/goconcurrency/workers"
	"log"
)

/*
The program fetches issues from the xkcd comics website and downloads each URL to build an offline JSON index.

credit to: https://blog.devgenius.io/concurrency-with-sample-project-in-golang-297400beb0a4

A goroutine is Go’s way of achieving concurrency. They are functions that run concurrently with other functions.
A goroutine can be compared to a lightweight thread (although it’s not a thread, as many goroutines can work on a single thread) which makes it lighter,
faster and reliable. You can create as many as one million goroutines in one program. When two or more goroutines are running, they need a way to communicate with each other.
That’s where channels come in.
*/

func main() {
	// allocate jobs
	noOfJobs := 3000
	go workers.AllocateJobs(noOfJobs)

	// get results
	done := make(chan bool)
	go workers.GetResults(done)

	// create worker pool
	noOfWorkers := 100
	workers.CreateWorkerPool(noOfWorkers)

	// wait for all results to be collected
	<-done

	// convert result collection to JSON
	data, err := json.MarshalIndent(workers.ResultCollection, "", "    ")
	if err != nil {
		log.Fatal("json err: ", err)
	}

	// write json data to file
	err = workers.WriteToFile(data)
	if err != nil {
		log.Fatal(err)
	}
}

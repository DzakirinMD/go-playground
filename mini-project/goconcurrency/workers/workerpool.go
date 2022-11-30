package workers

import (
	"fmt"
	"go_playground/mini-project/goconcurrency/api"
	"go_playground/mini-project/goconcurrency/structs"
	"log"
	"sync"
)

var jobs = make(chan structs.Job, 100)
var results = make(chan structs.Result, 100)
var ResultCollection []structs.Result

func AllocateJobs(noOfJobs int) {
	for i := 0; i <= noOfJobs; i++ {
		jobs <- structs.Job{i + 1}
	}
	close(jobs)
}

func GetResults(done chan bool) {
	for result := range results {
		if result.Num != 0 {
			fmt.Printf("Retrieving issue #%d\n", result.Num)
			ResultCollection = append(ResultCollection, result)
		}
	}
	done <- true
}

func CreateWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i <= noOfWorkers; i++ {
		wg.Add(1)
		go Worker(&wg)
	}
	wg.Wait()
	close(results)
}

func Worker(wg *sync.WaitGroup) {
	for job := range jobs {
		result, err := api.Fetch(job.Number)
		if err != nil {
			log.Printf("error in fetching: %v\n", err)
		}
		results <- *result
	}
	wg.Done()
}

package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, num)
		results <- num * num
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 5
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	//sending the jobs into the jobs channel
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	//Collect the results in an separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	//read from the results channel
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

}

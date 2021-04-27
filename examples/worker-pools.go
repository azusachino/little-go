package examples

import (
	"fmt"
	"time"
)

func worker_(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func init() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker_(w, jobs, results)
	}
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10 // Amount of work to process
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)
	// Launch goroutines to handle the work
	for gr := 0; gr < numberGoroutines; gr++ {
		go worker(tasks, gr)

	}
	// Add a bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}
	// Close the channel so the goroutines will quit when all the work is done
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			// This means the channel is empty and closed
			fmt.Printf("Worker: %d: Shutting Down\n", worker)
			return
		}
		fmt.Printf("Worker: %d: Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker: %d: Completed %s\n", worker, task)
	}
}

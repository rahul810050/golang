package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup){
	defer wg.Done() // one goroutine finished
	fmt.Printf("worker %d started\n", i)
	// some task is going on
	fmt.Printf("worker %d ended\n", i)
}

func main() {
	// fmt.Println("learning sync wait group in golang")

	var wg sync.WaitGroup
	// start 3 worker goroutines
	for i := 1; i <= 3; i++{
		wg.Add(1) // increment the waitgroup counter "it should be called before starting the goroutine"
		// waitgroup should be always passed by reference
		go worker(i, &wg) // calling goroutine 
	}




	/*
	it wont be executed sequencially even if we use waitgroup 
	because the goroutine runs independently it will call all the functions 
	whichever gets executed first will be printed first it is not in our hand 
	and also each time there will be different type of sequence in the output

	Q. why output order changes every run???
	Goroutines are sheduled by:
	- OS scheduler
	- Go runtime
	- CPU availability
	*/

	// wait for all workers to finish
	wg.Wait() // it will wait until the counter becomes zero
	fmt.Println("workers task complete")
}
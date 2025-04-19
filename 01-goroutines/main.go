package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Use the current time's nanoseconds to ensure a different seed each run
	rand.Seed(time.Now().UnixNano())

	begin := time.Now()
	fmt.Println("==== Construction begins ====")

	// Goroutines for concurrent workers
	go Worker("Bob", "installing a door")
	go Worker("Alice", "mowing the lawn")
	go Worker("John", "painting the walls")

	// Wait for goroutines to finish before ending the program
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("==== Construction ended in %.2f seconds ====\n", time.Since(begin).Seconds())
}

// Worker initiates a job and takes a random number of seconds between 3 and 5 to finish it
func Worker(name, job string) {
	fmt.Printf("Worker %s began to work on %s\n", name, job)

	for i := 1; i <= 10; i++ {
		fmt.Printf("Worker %s is working %d\n", name, i)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Printf("Worker %s finished working on %s\n", name, job)
}

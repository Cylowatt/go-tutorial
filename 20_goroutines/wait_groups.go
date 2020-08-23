package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func waitGroups() {
	// Add 1.
	wg.Add(1)

	go func() {
		sayHello()

		// Remove 1.
		wg.Done()
	}()

	// Wait until counter becomes 0.
	wg.Wait()
}

func wrongWaitGroupUsage() {
	for i := 0; i < 10; i++ {
		wg.Add(2)

		// Will print and increment in random order.
		// No synchronisation between these.
		go printCounter()
		go increment()
	}

	wg.Wait()
}

func printCounter() {
	fmt.Println("Counter:", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

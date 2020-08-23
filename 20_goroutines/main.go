package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// -1 does not change, and just returns
	// the current number of threads.
	fmt.Println("Threads:", runtime.GOMAXPROCS(-1))

	// This is a tuning variable.
	// Play with it to speed things up (performance testing).
	// Minimum: the amount of total CPU threads.
	runtime.GOMAXPROCS(24)
	fmt.Println("Threads:", runtime.GOMAXPROCS(-1))

	fmt.Printf("Blocking Invocations\n------------\n")
	blockingInvocations()

	fmt.Printf("\nWait Groups\n------------\n")
	waitGroups()
	wrongWaitGroupUsage()

	fmt.Printf("\nMutexes\n------------\n")
	mutexes()
}

func blockingInvocations() {
	go sayHello()

	msg := "Hello?"
	go func() {
		// Closures allows to capture the message.
		// Most likely will output "Goodbye", as this
		// get evaluated when the goroutine runs,
		// creating a race condition.
		//
		// Better way: pass an argument.
		fmt.Println(msg)
	}()
	msg = "Goodbye?"

	// Please never do this.
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}

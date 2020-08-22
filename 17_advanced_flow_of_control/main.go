package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("\nDefer Over Network\n------------\n")
	deferOverNetwork()

	fmt.Printf("\nDefer\n------------\n")
	deferPrint()

	fmt.Printf("\nPanic and Recover\n------------\n")
	panicRecover()
	fmt.Println("Recovered from panic")
}

func deferPrint() {
	a := "start"

	// Defer allows to move the execution to the end of the function.
	// However, deferred functions are executed in the reverse order
	// (LIFO - last in first out). This is to make sure that resources
	// are released in correct order.
	fmt.Println("First")
	defer fmt.Println("Second (deferred)")
	fmt.Println("Third")
	defer fmt.Println("Fourth (deferred)")
	fmt.Println("Fifth")

	// Arguments are passed at the time defer is called, so the result is "start".
	// Variable a is evaluated eagerly/captured.
	defer fmt.Println(a)
	a = "end"
}

func deferOverNetwork() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Common use case: associate opening and
	// closing a resource right next to each other.
	// Be careful with loops, as this is not run until the function exits.
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(robots))
}

func panicRecover() {
	// Return errors for expected cases.
	// Panic for something you cannot handle (behaviour like exceptions).

	// It's common to panic when a web server fails to start, for example.
	// Avoid using panic in libraries. Only top-level init code should.

	// This guy will run, because as the panic is propagated to the runtime,
	// the defer is still executed after the function exits.
	//
	// This is important to release resources even if the application panics.
	defer fmt.Println("This was deferred")

	// Defer takes a function not, and not the function itself, so you
	// need to execute the anonymous function. This will catch any
	// panic before this function exits.
	defer func() {
		// Recover from the panic and log the error.
		// The function will exit, but program execution will continue.
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// If you want to continue panicking:
			// panic(err)
		}
	}()

	panic("something bad happened")
	fmt.Println("Done with function - this will never be executed.")
}

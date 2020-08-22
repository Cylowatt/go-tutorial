package main

import "fmt"

func main() {
	fmt.Printf("Simple For Loop\n------------\n")
	simpleForLoop()

	fmt.Printf("\nFizz Buzz\n------------\n")
	fizzBuzz()

	fmt.Printf("\nMore For Loops\n------------\n")
	moreForLoops()
}

// A while loop is a for loop.
func simpleForLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	counter := 0
	for counter < 5 {
		fmt.Println(counter)
		counter++
	}
}

// Loop 1-100.
// For every multiple of 3 output fizz.
// For every multiple of 5 output buzz.
// For every multiple of 3 and 5 output fizzbuzz.
// For everything else, output the number.
func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if (i % 15 == 0) {
			fmt.Println("FizzBuzz")
		} else if (i % 3 == 0) {
			fmt.Println("Fizz")
		} else if (i % 5 == 0) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func moreForLoops() {
	// Multiple init.
	for i, j := 0, 0; i < 5; i, j = i + 1, j + 1 {
		fmt.Println("i:", i, "j:", j)
	}

	// Same as while(true).
	// continue is also supported.
	i := 0
	Outer:
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			fmt.Println("Break!")

			// Break out of the loop.
			// Labelling and breaking out by label also supported.
			break Outer
		}
	}
}
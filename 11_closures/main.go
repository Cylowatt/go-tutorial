package main

import "fmt"

func main() {
	// A closure is captured by declaring an anonymous function.
	sum := adder()

	for i := 1; i <= 10; i++ {
		fmt.Println(sum(i))
	}
}

// Returns a function that accepts an int and returns an int.
// The sum is a mutable part of the function, much like an object.
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}
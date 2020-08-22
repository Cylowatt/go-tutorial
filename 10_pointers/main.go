package main

import "fmt"

func main() {
	// A pointer allows you to point to a memory address of a value.
	// Useful for increasing performance when passing large values.
	// Also helpful to change values at specific locations.
	a := 5

	// b is a pointer to a.
	b := &a

	fmt.Println(a, b, *b)
	fmt.Printf("a: %T, b: %T\n", a, b)

	// Create a pointer and dereference it.
	fmt.Println(*&a)

	// Change the value of a through b.
	*b = 10
	fmt.Println(a)
}

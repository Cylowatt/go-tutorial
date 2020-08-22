package main

import "fmt"

type myStruct struct {
	foo int
	bar int
}

func main() {
	fmt.Printf("Struct Pointers\n------------\n")
	structPointers()
}

func structPointers() {
	// Using the new function. Cannot use initialisation syntax.
	// The result is a pointer.
	ms := new(myStruct)

	// The following two lines are the same.
	(*ms).foo = 50
	ms.foo = 50

	fmt.Println(ms)

	// Uninitiliased pointer will hold nil.
	// Check for nil in every function that takes pointers.
	var ms2 *myStruct
	fmt.Println(ms2)
}
package main

import (
	"fmt"
	"io"
	"reflect"
)

func getArea(s Shape) float64 {
	return s.area()
}

// Interfaces do not require explicit implementation, so you can create
// a wrapper interface that matches a real struct in a library,
// and you can use this struct as if it implements that interface.
//
// Any type can implement an interface, including your own type aliases
// to built-in Go primitives. Like extension methods in C#.
//
// Keep interfaces small (preferably one method) for maximum reusability.
// Compose interfaces to make larger interfaces.
//
// interface{} - everything is an empty interface.
// Useful for type-incompatible values and pointers.
// To use the underlying object, you will need
// to either use reflection or type conversion.
//
// Don't export interfaces for types that will be consumed.
// Only export the types. Let consumers deal with the interfaces.
// Do export interfaces for types that will be used by your package.
// Design functions and methods to receive interfaces whenever possible.
func main() {
	fmt.Printf("Shapes Example\n------------\n")
	runShapes()

	fmt.Printf("\nWriter Example\n------------\n")
	runWriter()

	fmt.Printf("\nIncrementer Example\n------------\n")
	runIncrementer()
}

func runShapes() {
	c := Circle{radius: 4}
	r := Rectangle{4, 4}

	fmt.Printf("Circle %v, area: %f\n", c, getArea(c))
	fmt.Printf("Rectangle %v, area: %f\n", r, getArea(r))
}

func runWriter() {
	cw := ConsoleWriter{}

	if _, err := cw.Write([]byte("Hello!")); err != nil {
		fmt.Println(err.Error())
	}
}

func runIncrementer() {
	// Same as below: var x IntCounter = 5
	x := IntCounter(5)
	var inc Incrementer = &x

	fmt.Println("Before increment:", x)
	fmt.Println("Increment:", inc.Increment())
	fmt.Println("After increment:", x)

	// Type conversion (downcast).
	// If conversion fails, the code will panic.
	ctr := inc.(*IntCounter)
	fmt.Println("Converted Incrementer interface to IntCounter", *ctr)

	// Error-driven type conversion.
	if rdr, ok := inc.(io.Reader); ok {
		fmt.Println("Converted to reader", rdr)
	} else {
		fmt.Println("Failed to convert to reader", reflect.TypeOf(inc))
	}
}

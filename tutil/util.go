package tutil

import "fmt"

// PV Prints a value and its type.
func PrintV(v interface{}) {
	fmt.Printf("%v, %T\n", v, v)
}

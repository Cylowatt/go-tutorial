package main

import "fmt"

func main() {
	arrays()
	slices()
	multiDimensionalArrays()
	sliceOperations()
}

// Arrays are of fixed length (compile-time) and of a given type.
func arrays() {
	var fruitArr [2]string

	fruitArr[0] = "mango"
	fruitArr[1] = "banana"

	fmt.Println(fruitArr)

	vegetableArr := [2]string{"broccoli", "cabbage"}
	fmt.Println(vegetableArr)

	// The ... allows the compiler to detect
	// the array size from the passed arguments.
	vegetableArr = [...]string{"broccoli", "cabbage"}
	fmt.Println(vegetableArr)
}

// Slices are dynamically allocated arrays.
// They are backed by arrays, but slices abstract that away.
func slices() {
	fruitSlice := []string{"apple", "orange", "papaya"}
	fmt.Println(fruitSlice, "slice value count", len(fruitSlice))

	fmt.Println("Slice capacity (underlying array size): ", cap(fruitSlice))

	// Slice [startIndex, endIndex).
	// [:] slice of all elements.
	// [:endIndex] same as [0:endIndex]
	// [startIndex:] same as [startIndex:len(slice) - 1]
	// Slicing works both with arrays and slices, but always returns slices.
	fmt.Println("[1:3] slice: ", fruitSlice[1:3])

	// Dynamically allocate a slice of arbitrary size and optional capacity.
	newSlice := make([]string, 10)
	fmt.Printf("Slice of size %d, %v\n", len(newSlice), newSlice)
}

func multiDimensionalArrays() {
	identityMatrix := [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	fmt.Println(identityMatrix)
}

func sliceOperations() {
	// Initialise empty slice.
	var newSlice []string

	// Add element to a slice. It may either create a new slice or use the same one.
	newSlice = append(newSlice, "elt1")
	fmt.Println(newSlice)

	// Spread an array/slice and append all of its elements via
	// the variadic append function using the ... operator.
	newSlice = append(newSlice, []string{"elt2", "elt3", "elt4", "elt5", "elt6"}...)
	fmt.Println(newSlice)

	// NOTE: all of the things below will screw up the underlying array.

	// Remove the first element from the slice.
	newSlice = newSlice[1:]
	fmt.Println(newSlice)

	// Remove the last element from the slice.
	newSlice = newSlice[:len(newSlice)-1]
	fmt.Println(newSlice)

	// Remove element at index 2.
	newSlice = append(newSlice[:2], newSlice[3:]...)
	fmt.Println(newSlice)
}

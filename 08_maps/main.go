package main

import "fmt"

// A map is a key-value pair.
func main() {
	fmt.Println("Map Creation\n-----------")
	mapCreation()
	fmt.Println("\nMap Mutation\n-----------")
	mapMutation()
}

func mapCreation() {
	// Inline.
	emails := map[string]string{"bob": "bob@bob.bob", "sally": "sally@sally.sally"}
	fmt.Println(emails)

	// Via make(type[, capacity]).
	numberNamesToNumbers := make(map[string]int)

	numberNamesToNumbers["one"] = 1
	numberNamesToNumbers["two"] = 2
	numberNamesToNumbers["three"] = 3

	fmt.Println(numberNamesToNumbers, numberNamesToNumbers["two"])
	printMapLength(numberNamesToNumbers)
}

func mapMutation() {
	numberNamesToNumbers := map[string]int{"one": 1, "two": 2, "three": 3}
	printMapLength(numberNamesToNumbers)

	numberNamesToNumbers["four"] = 4

	delete(numberNamesToNumbers, "three")
	fmt.Println(numberNamesToNumbers, numberNamesToNumbers["three"])
	fmt.Println("The map has", len(numberNamesToNumbers), "values")

	// Check if the key is present in the map, or is the value actually zero.
	// ok variable name is not a requirement, but is conventional.
	if num, ok := numberNamesToNumbers["six"]; ok {
		fmt.Printf("Value six is present. Value: %d\n", num)
	} else {
		fmt.Println("Value six is not present")
	}
}

func printMapLength(m map[string]int) {
	fmt.Println("The map has", len(m), "values")
}

package main

import "fmt"

// Range used to loop through arrays, maps, slices: collections.
func main() {
	fmt.Printf("\nRange Over Slice\n------------\n")
	rangeOverSlice()

	fmt.Printf("\nRange Over Map\n------------\n")
	rangeOverMap()

	fmt.Printf("\nString Ranges\n------------\n")
	stringRanges()

	fmt.Printf("\nChannel Ranges\n------------\n")
	channelRanges()
}

func rangeOverSlice() {
	ids := []int{34, 4, 54, 54, 4, 543, 3}

	for index, value := range ids {
		fmt.Printf("%d - ID: %d\n", index, value)
	}

	// Ignore the index.
	for _, value := range ids {
		fmt.Println(value)
	}
}

func rangeOverMap() {
	emails := map[string]string{"bob": "bob@bob.bob", "sally": "sally@sally.sally"}

	// Order not guaranteed.
	for key, email := range emails {
		fmt.Printf("name: %s, email: %s\n", key, email)
	}

	// Only use keys.
	for key := range emails {
		fmt.Println(key)
	}
}

func stringRanges() {
	s := "Hello 世界!"

	// Iterate over string runes.
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}

func channelRanges() {
	fmt.Println("NOT IMPLEMENTED")
}

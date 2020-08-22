package main

import "fmt"

func main() {
	fmt.Printf("Compare If\n------------\n")
	compareIf(5, 10)
	compareIf(10, 5)
	compareIf(5, 5)

	fmt.Printf("\nInit If\n------------\n")
	initIf()

	fmt.Printf("\nCompare Switch\n------------\n")
	compareSwitch("red")
	compareSwitch("blue")
	compareSwitch("green")
	compareSwitch("purple")

	fmt.Printf("\nInit Switch\n------------\n")
	initSwitch()

	fmt.Printf("\nTagless Switch\n------------\n")
	taglessSwitch(10)
	taglessSwitch(15)
	taglessSwitch(27)

	fmt.Printf("\nType Switch\n------------\n")
	typeSwitch(1)
	typeSwitch(1.0)
	typeSwitch("s")
	typeSwitch([0]int{})
}

func compareIf(x, y int) {
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is more than %d\n", x, y)
	}
}

func initIf() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// The num is scoped to the block.
	// What comes after ; is the condition for the block execution.
	if num, ok := m["one"]; ok {
		fmt.Println("one is present in the map as", num)
	}
}

func compareSwitch(colour string) {
	switch colour /* This value is called the tag. */ {
	case "green", "yellow": // Multiple tests.
		fmt.Println("color is green or yellow")
	case "red":
		fmt.Println("colour is red")
	case "blue":
		fmt.Println("colour is blue")
	default:
		fmt.Println("colour is not blue, red, green, nor yellow")
	}
}

func initSwitch() {
	// switch <init>; <tag> { ... }
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five, or ten")
	case 2, 4, 6:
		fmt.Println("two, four, or six")
	default:
		fmt.Println("another number")
	}
}

func taglessSwitch(i int) {
	// Same as if.
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")

		// Instead of break; like Swift.
		// Not too useful due to being able to supply
		// multiple test values for a tag switch.
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
}

func typeSwitch(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break // Leave the case early if there is some logical test.
		fmt.Println("Some additional stuff")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}

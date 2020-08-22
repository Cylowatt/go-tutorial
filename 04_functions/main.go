package main

import "fmt"

func main() {
	fmt.Println(greeting("Len"))
	fmt.Println(sum(4, 6))
	fmt.Println(sumInferType(4, 6))
}

func greeting(name string) string {
	return "Hello " + name
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func sumInferType(num1, num2 int) int {
	return num1 + num2
}
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

// we can return 2 or more variables
func swap(a, b int) (int, int) {
	return b, a
}

// here result is the returning variable so no need to mention it while returning
func multiply(a, b int) (result int) {
	// there is no need to declare result when it is mentioned as the returning variable
	result = a * b
	return
}

func main() {
	fmt.Println(add(1, 5));

	a, b := 1, 2
	fmt.Println("a = ", a, " b = ", b)

	a, b = swap(a, b)
	fmt.Println("a = ", a, " b = ", b)

	fmt.Println(multiply(5, 7))
}
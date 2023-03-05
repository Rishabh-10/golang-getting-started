package main

import "fmt"

func add(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println("Sum: ", add(1,2,3,4,5))

	arr := []int{1,2,3}
	fmt.Println("Sum: ", add(arr...))
}
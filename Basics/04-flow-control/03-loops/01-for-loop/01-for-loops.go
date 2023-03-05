package main

import "fmt"

func main() {
	sum := 0
	// simple for loop
	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println("Using simple for loop:", sum)

	// while loop using for
	sum = 0
	var iter int = 1
	for iter <= 10 {
		sum += iter
		iter++
	}
	fmt.Println("Using forWhile loop:", sum)

	// infinte loop using for
	iter = 1
	for {
		if iter%10 == 0 {
			break;
		}
		iter++
	}
	fmt.Println("Using infinte loop:", iter)

	fmt.Println("\nPrinting Odd numbers till 10: ")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
}
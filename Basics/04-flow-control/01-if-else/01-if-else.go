package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = 4

	if i%2 ==0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// the scope of the variable num will be limited in the if-else block
	// not outside the if else block
	if num := math.Pow(3, 2); num < 5 {
		fmt.Println(5)
	} else {
		fmt.Println(num)
	}
}
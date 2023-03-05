package main

import "fmt"

func main() {
	var a[5]int

	// zero value of array depend on the data type it will be containing
	fmt.Println("Array zero value:", a)

	fmt.Println("Lenght:", len(a), "\nCapacity:", cap(a))

	b := [5]int{1,2,3,4,5}
	fmt.Print("\nElements inside array: ")
	for _, v := range b {
		fmt.Print(v, " ")
	} 

	fmt.Println()

	// we can print the array without traversing it as well
	fmt.Println("Array without traversing:", b)
}
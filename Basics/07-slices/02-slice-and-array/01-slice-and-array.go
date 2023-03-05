package main

import "fmt"

func main() {
	a := [5]int{1,2,3,4,5}

	fmt.Println("Array a:", a)

	// making a slice from array
	// the below slice will contain all elements of the array
	sli := a[:]
	fmt.Println("Slice sli:", sli)
	// a slice consists of 3 things an address to an array
	// length, the number of elements inside it
	// capacity, the size of the array it is pointing to
}
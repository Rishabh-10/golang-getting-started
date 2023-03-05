package main

import "fmt"

func main() {
	// to declare a slice we need to keep the size field empty
	var s = []int{1, 2, 3}
	fmt.Println(s)


	// slices are dynamic array
	// make(slice, len, capacity)
	a := make([]int, 5, 10)

	fmt.Println(a)
	fmt.Println("Lenght:", len(a))
	fmt.Println("Capacity:", cap(a))

	// appending in a slice
	a = append(a, 10, 20)

	fmt.Println("After appending:", a)

	// appending a slice to a slice
	s = append(s, a...)
	fmt.Println("After appending slice a to s:", s)

	// copy a slice to another
	// while copying we need a slice with more than 0 length
	// otherwise nothing will get copied
	copySlice := make([]int, 5)
	copy(copySlice, s)

	// we need to create an empty slice c of the same length as s to copy into c from s.
	// copy(c, s)

	fmt.Println(copySlice)
}
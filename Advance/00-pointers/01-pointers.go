package main

import "fmt"

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.
func byVal(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.
func byRef(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	byVal(i)
	fmt.Println("After passing to byVal:", i)

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.
	byRef(&i)
	fmt.Println("After passing to byRef:", i)

	// Pointers can be printed too.
	fmt.Println("Address of i:", &i)
}
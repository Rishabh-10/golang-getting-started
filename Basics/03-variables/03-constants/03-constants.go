package main

import "fmt"

const s string = "constant"

func main() {

	const n = 3.14
	fmt.Println("s =", s)
	fmt.Println("n =", n)

	fmt.Printf("Type of s = %T", s)
	fmt.Printf("\nType of n = %T", n)

	anyFunction()
}

func anyFunction() {
	// as we can see that s was a global variable
	// but still we can reinitialize it in other functions
	// This is because the compiler gives preference to local variables than global. 
	// Hence, if a local and global variable share a name, the compiler will always choose the local variable.

	s := "Hello"

	fmt.Println("\nGlobal Constants: ", s)
}
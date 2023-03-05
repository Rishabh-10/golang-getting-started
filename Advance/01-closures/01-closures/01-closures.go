package main

import "fmt"


// This function intSeq returns another function, 
// which we define anonymously in the body of intSeq.  
// The returned function closes over the variable i to  
// form a closure.

// The function inside intSeq will be called anonymous function.

// An 'anonymous function' (function literal, lambda 
// abstraction, lambda function, lambda expression or 
// block) is a function definition that is not bound to 
// an identifier. Anonymous functions are often 
// arguments being passed to higher-order functions or 
// used for constructing the result of a higher-order 
// function that needs to return a function.


// How do closures work??????
// In Golang, a closure is a function that references variables outside of its scope. A closure can outlive the scope in which it was created. Thus it can access variables within that scope, even after the scope is destroyed.
func intSeq() func () int {
	i := 0
	return func () int {
		i++
		return i
	}
}

func main() {
	
	next := intSeq()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	fmt.Println("\nNew variable created")

	// here when we create a new variable/closure then this will create a new ref
	// to the variable i as in intSeq will run again an i will we reinitialized
	// but it will not disturb the i that is connected to next(previous variable)
	next2 := intSeq()
	fmt.Println(next2())
	
	fmt.Println("\nAccessing the previous variable")
	fmt.Println(next())
}
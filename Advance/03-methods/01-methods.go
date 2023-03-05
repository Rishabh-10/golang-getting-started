package main

import (
	"fmt"
	"net/http"
)

type Pair struct {
	X, Y int
}

// Methods can be defined for either pointer or value receiver types. 

// Here's an example of a value receiver
func (p Pair) add() int {
	return p.X + p.Y
}

// Here's an example of a pointer receiver
func (p *Pair) multiply() int {
	return p.X * p.Y
}


// Go automatically handles conversion between values and pointers for method calls. 
// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct

func main() {
	p1 := Pair{5, 10}

	fmt.Println("Addition: ", p1.add())
	fmt.Println("Multiplication: ", p1.multiply())

	p2 := &Pair{15, 25}
	fmt.Println("\nAddition: ", (*p2).add())
	fmt.Println("Multiplication: ", p2.multiply())
}	
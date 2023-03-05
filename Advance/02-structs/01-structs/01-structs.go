package main

import "fmt"

type Pair struct {
	X, Y int
}

func main() {
	var p1 Pair

	// the zero value is the zero value of its field here it will be {0, 0}
	fmt.Printf("Zero Value: %v\n", p1)

	p1 = Pair{1, 2}
	fmt.Println("p1:", p1)

	p2 := Pair{3, 4}
	fmt.Println("p2:", p2)

	var p3 = Pair{5, 6}
	fmt.Println("p3:", p3)

	var p4 = &p3
	fmt.Println("p4:", p4)

	// Golang implicitly dereferences so it is optional to use (*p4).X
	p4.X = 10
	fmt.Println("p3 after changing X in p4:", p3)

	(*p4).Y = 20
	fmt.Println("p3 after changing Y in p4:", p3)


	p5 := Pair{
		X: 50,
		Y: 100,
	}

	fmt.Println("p5:",p5)
}
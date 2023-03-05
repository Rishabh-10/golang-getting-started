package main

import "fmt"

func main() {
	var x int
	x = 0

	var a int = 5
	var b, c int = 10 , 15


	fmt.Printf("x = %v\na = %v\nb = %v\nc = %v\n\n", x, a, b, c)
	fmt.Println(x, a, b, c)
	fmt.Print("\n", a, b, c)
}
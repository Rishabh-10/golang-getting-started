package main

import "fmt"

func main() {
	pair1 := struct {
		X int
		Y int
	}{
		X: 1,
		Y: 2,
	}

	fmt.Println(pair1, pair1.X, pair1.Y)
}
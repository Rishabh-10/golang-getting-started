package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// seed is used so that we get different numbers every time
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random Number: ", rand.Intn(20))
	fmt.Println("Random Number: ", rand.Intn(20))

	fmt.Println("\nPI: ", math.Pi)

	fmt.Println("\nSqrt of 25: ", math.Sqrt(25))
}
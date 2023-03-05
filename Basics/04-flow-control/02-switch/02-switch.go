package main

import (
	"fmt"
	"time"
)

func main() {
	num := 2

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("No Idea")
	}

	// as in the below switch case we are looking for a true case
	// so we can mention true in the expression field
	// or leave it empty
	n := 2
	switch true {
	case n < 5:
		fmt.Println("Less than 5")
	case n == 5:
		fmt.Println("Equal to 5")
	case n > 5:
		fmt.Println("More than 5")
	}

	// here switch can take strings as input as well
	input := time.Now()
	switch {
	case input.Hour() < 12:
		fmt.Println("Good Morning")
	case input.Hour() >= 12 && input.Hour() < 18:
		fmt.Println("Good Afternoon")
	case input.Hour() >= 18 && input.Hour() <= 21:
		fmt.Println("Good Evening")
	case input.Hour() > 21 && input.Hour() <= 3:
		fmt.Println("Good Night")
	}
}
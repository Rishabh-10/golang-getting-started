package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()

	input := "One Piece"
	// for index, value := range dataStructure{}
	for _, v := range input {
		// here v will be ASCII code so we need to type cast v
		fmt.Print(string(v), " ")
	}
	
	// IMP
	// range on strings iterates over Unicode code points.
}
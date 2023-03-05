package main

import "fmt"

func main() {
	// make(map[key-type]val-type)
	mp := make(map[string]int)

	arr := [8]string{"Hi", "Hello", "Hi", "Close", "Hello", "Close", "Bye", "Close"}
	fmt.Println("Array arr:", arr)
	for _, val := range arr {
		mp[val] += 1
	}

	fmt.Println("\nMap mp:", mp)
	fmt.Println()
	for key, val := range mp {
		fmt.Println(key, "-", val)
	}

	delete(mp, "Close")
	fmt.Println("\nMap mp:", mp)
	fmt.Println()
	// to check if a key exists
	// val, flag := mp['keyName']
	// flag returns true or false
	// val returns the value if it does not exists then zero value is returned
	val, flag := mp["Close"]
	
	fmt.Println(val)
	if flag == false {
		fmt.Println("Close does not exists")
	} else {
		fmt.Println("Close exists")
	}
}
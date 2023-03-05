package main

import "fmt"

func recursiveClosure() {
	i := 0

	var manyGodByes func ()

	manyGodByes = func () {
		if i < 5 {
			fmt.Println("Bye!!!")
		} else {
			return
		}
		i++
		manyGodByes()
	}

	manyGodByes()
}

func main() {
	i := 0

	var manyGreeting func ()

	manyGreeting = func () {
		if i < 5 {
			fmt.Println("Hello!!!")
		} else {
			return
		}
		i++
		manyGreeting()
	}

	manyGreeting()

	recursiveClosure()
}
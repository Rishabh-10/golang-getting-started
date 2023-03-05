package main
  
import "fmt"
  
func main() {
      
    // Declaring the variable
    GFG := 0
  
    // Assigning an anonymous  
    // function to a variable 
    counter := func() int {
       GFG += 1
       return GFG
    }

	counter2 := func() int {
		GFG += 5
		return GFG
	}

	// here both the variables will be referencing to the same variable
	// so both will affect the variable GFG
  
    fmt.Println(counter())
    fmt.Println(counter2())
	fmt.Println(counter())
}

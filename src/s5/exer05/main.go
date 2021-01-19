// customer type - exercise 5
package main

import "fmt"

// X Package level type
type animal int

var x animal

var y int

func main() {

	fmt.Printf("Value of x = %v\n", x)
	fmt.Printf("type of x is %T\n", x)

	x = 55
	fmt.Printf("value of x is %d\n", x)

	y = int(x)
	fmt.Printf("value of y is %d\n", y)
	fmt.Printf("Type of y is %T\n", y)
}

// customer type - Exercise 4
package main

import "fmt"

type myType int

func main() {

	var x myType

	fmt.Printf("Value of x = %d\n", x)
	fmt.Printf("Type of x = %T\n", x)

	x = 10
	fmt.Printf("x = %d, x is of type %T\n", x, x)
}

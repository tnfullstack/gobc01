// customer type - Exercise 4
package main

import "fmt"

type myType int

func main() {

	var x myType
	x = 10

	fmt.Printf("x = %d, x is of type %T\n", x, x)
}

// Section 5 - exercise 2
package main

import "fmt"

// package level variable
var (
	X int
	Y string
	Z bool
)

func main() {
	fmt.Println("Zero value for int =", X)
	fmt.Println("Zero value for string =", Y)
	fmt.Println("Zero value for bool =", Z)
}

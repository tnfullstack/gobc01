// Section 5 - exercise 3
package main

import "fmt"

// Exported variables
var (
	X int
	Y string
	Z bool
)

func main() {
	X = 42
	Y = "James Bond"
	Z = true

	s := "Something to test short declaration of type string"
	fmt.Printf("s variable hold '%s'\n", s)

	a := fmt.Sprintf("%d %s %v\n", X, Y, Z)
	fmt.Println(a)
}

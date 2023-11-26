// Standard input with Scanf
package main

import (
	"fmt"
)

func main() {
	var (
		name string
		age  int
	)
	fmt.Print("What is your name? ")
	fmt.Scanf("%s\n", &name)

	fmt.Print("How old are you? ")
	fmt.Scanf("%d\n", &age)
	fmt.Printf("%s is %d years old!\n", name, age)
}

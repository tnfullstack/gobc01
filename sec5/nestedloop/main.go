// Nested loop
package main

import "fmt"

func main() {

	// Outer loop
	for i := 1; i <= 10; i++ {
		fmt.Println("Outter loop", i)
		for j := 1; j <= 5; j++ {
			fmt.Println("          => Inner loop", j)
		}

	}
}

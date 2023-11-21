// For loop
package main

import "fmt"

func main() {
	// Declare a variable num and assing interger 100
	num := 100
	sum := 0

	// Adding 1 to 100 and print the all values using for loop
	for i := 1; i <= num; i++ {
		fmt.Printf("#%d + %d = ", i, sum)
		sum += i
		fmt.Printf("%d\n", sum)
	}
}

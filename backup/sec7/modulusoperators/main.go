// Modulus operators
package main

import "fmt"

const (
	ROCK     = 0 // beats scissors (SCISSORS + 1) % 3 = 0
	PAPER    = 1 // beats rock (ROCK + 1) % 3 = 1
	SCISSORS = 2 // beats paper(PAPER + 1) % 3 = 2
)

func main() {

	// modulus operators
	a := (SCISSORS + 1) % 3 // => 0
	fmt.Println("(2 + 1) % 3 =", a)
	b := (ROCK + 1) % 3 // => 1
	fmt.Println("(0 + 1) % 3 =", b)
	c := (PAPER + 1) % 3 // => 2
	fmt.Println("(1 + 1) % 3 =", c)

}

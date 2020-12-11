// Randomization and Go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	guess := 9

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		if n == guess {
			fmt.Printf("You Win!, your guess = %d, rand %d\n", guess, n)
			return
		}
		fmt.Printf("You loss rand %d\n", n)
	}
	fmt.Println()
}

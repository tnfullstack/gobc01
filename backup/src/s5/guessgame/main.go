// Guessing game
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	const maxTurn = 5

	t := time.Now().UnixNano()
	rand.Seed(t)

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please enter your guess number.")
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.", err)
		return
	}

	if guess < 0 {
		fmt.Println("Please enter a positive number.")
		return
	}

	for cguess := 0; cguess < maxTurn; cguess++ {
		cguess = rand.Intn(guess + 1)

		if guess == cguess {
			fmt.Printf("🏆 YOU WIN!, your guess %d, computer guess %d\n", guess, cguess)
			return
		}
	}
	fmt.Println("💀 You loss!")
}

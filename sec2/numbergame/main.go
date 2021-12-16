// Guess a number game
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	enter = "then press ENTER when ready"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	firstNum := rand.Intn(8) + 2
	secondNum := rand.Intn(8) + 2
	subtraction := rand.Intn(8) + 2
	var answer int

	reader := bufio.NewReader(os.Stdin)

	// Display a game instruction
	fmt.Println("")
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", enter)
	fmt.Print("-> ")
	reader.ReadString('\n')

	// Walk user through the game
	fmt.Println("Multiply your number by", firstNum, enter)
	fmt.Print("-> ")
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNum, enter)
	fmt.Print("-> ")
	reader.ReadString('\n')

	fmt.Println("Devide the result by the original number", enter)
	fmt.Print("-> ")
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, enter)
	fmt.Print("-> ")
	reader.ReadString('\n')

	// Display the answer
	answer = (firstNum * secondNum) - subtraction
	fmt.Println("The answer is", answer)

}

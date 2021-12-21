// Rock Paper Scissor Game
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {

	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(3)

	// Clear the screen each time the game start
	clearScreen()

	// Set the number of time the game repeat play
	round := 3

	// The for loop allow the game play upto the number of round defined above
	for round >= 0 {
		// Read keyboard input from command-line terminal
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Please enter rock, paper, or scissors -> ")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		playerChoice = strings.ToLower(playerChoice)

		// if and else if statements
		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		}

		// switch statements practice
		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")
		case PAPER:
			fmt.Println("Computer chose PAPPER")
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")
		default:
		}

		// Print out the user input choise
		fmt.Println("Player choice", strings.ToUpper(playerChoice))
		fmt.Println()

		// Decide the winner
		if playerValue == computerValue {
			fmt.Println("It's a tide.")
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					fmt.Println("Computer wins.")
				} else {
					fmt.Println("Player wins.")
				}
			case PAPER:
				if computerValue == SCISSORS {
					fmt.Println("Computer wins.")
				} else {
					fmt.Println("Player wins.")
				}
			case SCISSORS:
				if computerValue == ROCK {
					fmt.Println("Computer wins.")
				} else {
					fmt.Println("Player wins.")
				}
			default:
				fmt.Println(playerChoice, "is not and option.")
			}
		}

		// game round countdown, the the round cound down to zero - return (stop program)
		round--
		if round == 0 {
			return
		}
	}
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

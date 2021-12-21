// Rock | Paper | Scissors - Game Logic
package game

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

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {

	// Process input using select channels

	// Print Game information on screen

	// Keep track of game information
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
		}
	}

}

// ClearScreen clears the screen
func (g *Game) ClearScreen() {
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

// Print game intro
func (p *Game) PrintInfo() {
	fmt.Println("Rock, Paper, Scissors")
	fmt.Println("---------------------")
	fmt.Println("The game will play 3 rounds, best of three wins.")
	fmt.Println()
}

// PlayRound start the game rounds
func (g *Game) PlayRound() bool {

	rand.Seed(time.Now().UnixNano())

	playerValue := -1

	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumber)
	fmt.Println("------")

	fmt.Print("Please enter Rock, Paper, or Scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	} else {
		playerValue = -1
	}

	fmt.Println()
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK")
	case PAPER:
		fmt.Println("Computer chose PAPER")
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS")
	default:
	}

	if playerValue == computerValue {
		g.DisplayChan <- "It is a draw!"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
		default:
			g.DisplayChan <- "Invalid choice!"
			return false
		}
	}
	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
}

func (g *Game) PrintFinalScore() {
	fmt.Println()
	fmt.Println("Best of three result:")
	fmt.Println("Computer score:", g.Round.ComputerScore)
	fmt.Println("Player score:", g.Round.PlayerScore)
	if g.Round.ComputerScore < g.Round.PlayerScore {
		fmt.Println("Player wins!")
	} else {
		fmt.Println("Computer wins!")
	}
}

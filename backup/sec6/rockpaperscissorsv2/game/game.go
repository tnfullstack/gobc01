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

//
func (g *Game) Rounds() {

	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 0
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			g.DisplayChan <- ""
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
func (p *Game) PrintIntro() {
	p.DisplayChan <- `
Rock, Paper, Scissors
----------------------
The game will play 3 rounds, best of three wins.

`
	<-p.DisplayChan
}

// PlayRound start the game rounds
func (g *Game) PlayRound() bool {

	rand.Seed(time.Now().UnixNano())

	playerValue := -1

	g.DisplayChan <- fmt.Sprintf(`
Round %d
--------
	`, g.Round.RoundNumber)
	<-g.DisplayChan

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
	}

	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<-g.DisplayChan

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose ROCK"
		<-g.DisplayChan
	case PAPER:
		g.DisplayChan <- "Computer chose PAPER"
		<-g.DisplayChan
	case SCISSORS:
		g.DisplayChan <- "Computer chose SCISSORS"
		<-g.DisplayChan
	default:
	}

	if playerValue == computerValue {
		g.DisplayChan <- "It is a draw!"
		<-g.DisplayChan
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
			g.DisplayChan <- fmt.Sprintf("%s is an anvalid choice!", playerChoice)
			<-g.DisplayChan
			return false
		}
	}
	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
	<-g.DisplayChan
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
	<-g.DisplayChan
}

func (g *Game) PrintFinalScore() {

	g.DisplayChan <- fmt.Sprintf(`
Best of three results:
Computer score: %d
Player score: %d
`, g.Round.ComputerScore, g.Round.PlayerScore)
	<-g.DisplayChan

	if g.Round.ComputerScore < g.Round.PlayerScore {
		g.DisplayChan <- "Player wins!"
		<-g.DisplayChan
	} else {
		g.DisplayChan <- "Computer wins!"
		<-g.DisplayChan
	}
}

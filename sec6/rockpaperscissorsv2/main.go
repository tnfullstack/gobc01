// Rock | Paper | scissors game using channel
package main

import (
	"go/gobc01/sec6/rockpaperscissorsv2/game"
)

func main() {

	// Declare and initialize chan
	displayChan := make(chan string)
	roundChan := make(chan int)

	// game stores channel and game.Round data
	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	// go routine
	go game.Rounds()

	// ClearScress clears terminal before each game
	game.ClearScreen()

	// Print game info
	game.PrintInfo()

	for {
		game.RoundChan <- 1
		<-game.RoundChan

		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}

	// Print final score
	game.PrintFinalScore()

}

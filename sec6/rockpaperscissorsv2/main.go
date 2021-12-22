// Rock | Paper | scissors game using channel
package main

import (
	"go/gobc01/sec6/rockpaperscissorsv2/game"
)

func main() {

	// Declare and initialize chan
	displayReCh := make(chan string)
	roundReCh := make(chan int)

	// game stores game.struct and game.Round data
	game := game.Game{
		DisplayChan: displayReCh,
		RoundChan:   roundReCh,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	// go routine - call Round function from game package
	go game.Rounds()

	// ClearScress clears terminal before each game
	game.ClearScreen()

	// Print game info
	game.PrintIntro()

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

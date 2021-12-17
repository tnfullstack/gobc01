// Hammer bitcoin game
package main

import (
	"fmt"

	"github.com/tnfullstack/gobc01/sec3/hammerbitcoin/game"
)

func main() {
	playAgain := true

	for playAgain {
		game.Play()
		playAgain = game.GetYesOrNo("Would you like to play again (y/n)?")
	}

	fmt.Println("")
	fmt.Println("Goodbye.")
}

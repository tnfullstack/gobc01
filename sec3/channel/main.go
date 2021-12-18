// Channel
package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPress chan rune

func main() {
	keyPress = make(chan rune)

	go listenForKeypress()

	fmt.Println("Press any key or q to quit.")

	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		//
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}
		keyPress <- char
		// fmt.Println("Test", keyPress, &char)
	}
}

func listenForKeypress() {
	for {
		key := <-keyPress
		fmt.Println("You press", key, string(key))
	}
}

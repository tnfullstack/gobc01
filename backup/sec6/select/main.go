// Select - channel
package main

import (
	"fmt"
	"time"
)

// Create channels
var chan1 = make(chan string)
var chan2 = make(chan string)

func task1() {
	time.Sleep(time.Second * 1)
	chan1 <- "One"
}

func task2() {
	time.Sleep(time.Second * 2)
	chan2 <- "Two"
}

func main() {

	// Go routine
	go task1()
	go task2()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received", msg2)
		}
	}

}

package logger

import "log"

func ListenForLog(ch chan string) {
	for {
		// receive message from channel ch
		msg := <-ch

		log.Println("You have enter =>", msg)
	}
}

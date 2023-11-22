// Shutting down the application gracefully
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var writer *os.File

func main() {
	// The file is opened as a log file to write into.
	// This way we represent the resources allocation.
	file := fmt.Sprintf("test_%d.log", time.Now().Unix())
	writer, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// The code is running in a goroutine independently. So in case the
	// program is terminated from outside, we need to let the goroutine
	// know via the closeChan
	closeChan := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-closeChan:
				log.Println("Goroutine closing")
				return
			default:
				log.Println("Writing to log")
				io.WriteString(writer, fmt.Sprintf("Logging access %s\n", time.Now()))
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT)

	// This is blocking read from sigChan where the Notify function sends the signal.
	<-sigChan

	// After the signal is received all the code behind the read from channel could
	// be considered as a cleanup.
	// CLEANUP SECTION
	close(closeChan)
	releaseAllResources()
	fmt.Println("The application shutdown gracefully")
}

func releaseAllResources() {
	io.WriteString(writer, "Application releasing all resources\n")
	writer.Close()
}

// Creating temporary files
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tFile, err := os.CreateTemp("", "temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tFile.Name()) // clean up

	fmt.Println(tFile.Name())

	// TempDir returns the path in string
	tDir := os.TempDir()

	defer os.Remove(tDir)
	fmt.Println(tDir)
}

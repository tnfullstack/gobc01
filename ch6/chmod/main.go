// Changing file permissions
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("test.file")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Obtain current permissions
	fStat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File name: %s\n", fStat.Name())
	fmt.Printf("File permissions: %v\n", fStat.Mode())

	// Change file permissions
	err = f.Chmod(0777)
	if err != nil {
		log.Fatal(err)
	}
	fNewStat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New permissions %v\n", fNewStat.Mode())
}

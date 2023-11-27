// Getting file information
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	file, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %s\n", file.Name())
	fmt.Printf("Is directory: %t\n", file.IsDir())
	fmt.Printf("Size: %d\n", file.Size())
	fmt.Printf("Mode: %v\n", file.Mode())
}

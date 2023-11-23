// Concatenating a string with writer
package main

import (
	"bytes"
	"fmt"
)

func main() {
	strings := []string{"This ", "is ", "even ", "more ", "performant "}

	// declare a bytes buffer to store string
	buffer := bytes.Buffer{}
	for _, val := range strings {
		// iterates through the string slice and store each string
		// to buffer using -> buffer.WriteString(s string) (int, error)
		buffer.WriteString(val)
	}

	// Print the buffer's data to the screen
	fmt.Println(buffer.String())
}

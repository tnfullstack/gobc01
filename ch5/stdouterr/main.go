// Writing standout output and error
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	testStr := "This is string to standard output.\n"
	testStr2 := "This is string to standard error.\n"

	// Simple write string
	io.WriteString(os.Stdout, testStr)

	io.WriteString(os.Stderr, testStr2)

	// Stdout/err implements writer interface
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, err := os.Stdout.Write(buf); err != nil {
			panic(err)
		}
	}

	// The fmt package could be used too
	fmt.Fprint(os.Stdout, "\n")
}

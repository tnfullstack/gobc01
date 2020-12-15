// bufio.scanner Scan an input stream line by line into a buffer
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// os.Stdin.Close()

	// Setting default scanning line by line
	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		lines++
	}

	fmt.Println("The number of lines counted:", lines)

	// Checking for file err
	if err := in.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

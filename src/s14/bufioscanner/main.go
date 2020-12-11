// bufioscanner - example on how bufio.scanner works
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	var line int

	for in.Scan() {
		line++
	}
	os.Stdin.Close()
	fmt.Printf("Number of lines: %d\n", line)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}

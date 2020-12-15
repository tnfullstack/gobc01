// Use maps as sets
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: command < [filename]")
		return
	}

	search := args[0]

	// default scan from stendard input
	in := bufio.NewScanner(os.Stdin)

	// setup scan for words from standard input
	in.Split(bufio.ScanWords)

	// Setup a words map variable
	words := make(map[string]bool)

	for in.Scan() {

		word := strings.ToLower(in.Text())

		if len(word) > 2 {
			words[word] = true
		}
	}

	if words[search] {
		fmt.Printf("The input contains %q.\n", search)
		return
	}
	fmt.Printf("The input does not contains %q.\n", search)

}

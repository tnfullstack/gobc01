// mapsets -
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please enter a search word.")
		return
	}

	rx := regexp.MustCompile(`[^a-z]+`)

	query := args[0]

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")
		if len(word) > 2 {
			words[word] = true
		}
	}

	/*
		// Print words map for testing purpse only
		for k := range words {
			fmt.Println(k)
		}
	*/

	// query := "sun"
	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}
	fmt.Printf("Sorry, the input does not contain %q\n", query)

}

// uniquewords -
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
//
// ---------------------------------------------------------

func main() {
	var wc int

	// remove all charactors except for letter a-z
	rx := regexp.MustCompile(`[^a-z]+`)

	// Reading file using bufio.NewScanner(os.Stdin())
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	// Create a map to hold unique word
	words := make(map[string]bool)

	// cycle throught the bufio and compute each word
	for in.Scan() {
		wc++
		word := in.Text()

		// Replace any charactor that are not a-z with empty string ""
		word = rx.ReplaceAllString(word, "")
		fmt.Printf("%-5d %s\n", wc, word)

		// check if word exist
		if words[word] {
			fmt.Printf("Word already in map %q\n", word)
		}
		// add word in map words[string]bool
		words[word] = true
	}

	// Print out the total and unique word counts
	fmt.Printf("Total word count %d, number of unique word %d\n", wc, len(words))
}

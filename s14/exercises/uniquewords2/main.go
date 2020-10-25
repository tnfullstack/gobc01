package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func main() {

	var (
		wc int
	)
	// This is the regular expression pattern you need to use:
	// [^A-Za-z]+
	//
	// Matches to any character but upper case and lower case letters
	rx := regexp.MustCompile(`[^A-Za-z]+`)

	// Read from stdin and save data in bufio
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	// Declare words map to hold all unique words
	var words = make(map[string]bool)

	for in.Scan() {
		// word count
		wc++

		// convert any upper case leter to lowercase
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")

		// Add word to words map
		words[word] = true

		// Print exmple of uniqe words (for testing only)
		// fmt.Printf("%-10d %s\n", wc, word)
	}
	fmt.Printf("There are %d words, %d of them are unique.\n", wc, len(words))
}

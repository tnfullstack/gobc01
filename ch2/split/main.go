// Splitting string on an instance of substring
package main

import (
	"fmt"
	"strings"
)

func main() {

	const testStr1 = "Mary_had a little_lamb"
	const testStr2 = "Mayy*has,a%little_lamb!duck"
	const specialChar = "*,%_!"

	words := strings.Split(testStr1, "_")
	for i, substr := range words {
		fmt.Printf("%d: %s\n", i, substr)
	}

	splitFunc := func(r rune) bool {
		return strings.ContainsRune(specialChar, r)
	}
	subStr := strings.FieldsFunc(testStr2, splitFunc)

	for i, word := range subStr {
		fmt.Printf("Word %d is: %s\n", i, word)
	}
}

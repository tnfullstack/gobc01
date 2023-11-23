// Breaking string into words
package main

import (
	"fmt"
	"strings"
)

func main() {
	const testStr = "Mary has a little lamb"

	// split string when one or more consecutive of white space occur
	words := strings.Fields(testStr)
	for i, word := range words {
		fmt.Printf("%d: %s\n", i, word)
	}
}

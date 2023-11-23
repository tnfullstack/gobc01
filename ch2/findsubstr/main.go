// Find the substring in string
package main

import (
	"fmt"
	"strings"
)

func main() {
	const str = "Mary had a little lamb"

	subStr := "lamb"

	trueOrFalse := strings.Contains(str, subStr)
	fmt.Printf("The %s, contains %q: %t \n", str, subStr, trueOrFalse)

	subStr = "wolf"
	trueOrFalse = strings.Contains(str, subStr)
	fmt.Printf("The %s, contains %q: %t\n", str, subStr, trueOrFalse)

	startStr := "Mary"
	trueOrFalse = strings.HasPrefix(str, startStr)
	fmt.Printf("The %s, starts with %q: %t\n", str, startStr, trueOrFalse)

	endStr := "lamb"
	trueOrFalse = strings.HasSuffix(str, endStr)
	fmt.Printf("The %s, end with %q: %t\n", str, endStr, trueOrFalse)
}

// Replace string using strings.NewReplacer(string)
package main

import (
	"fmt"
	"strings"
)

func main() {
	const testStr = "Mary has a little lamb"

	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack", "little", "big")
	out := replacer.Replace(testStr)
	fmt.Println(out)
}

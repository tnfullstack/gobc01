// Replacing part of the string
package main

import (
	"fmt"
	"strings"
)

func main() {
	const testStr = "Mary has a little lamb"
	const testStr2 = "lamb lamb lamb lamb"

	out := strings.Replace(testStr, "little", "wolf", -1)
	fmt.Println(out)

	out = strings.Replace(testStr2, "lamb", "wolf", 2)
	fmt.Println(out)
}

// How to decode a string?
package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	word := []byte("phố huế")

	fmt.Printf("%s = % [1]x\n", word)

	/*
		var size int
		for i, r := range string(word) {
			fmt.Printf("%d %c\n", i, r)

			if i > 0 {
				size = i
				break
			}
		}
	*/

	r, size := utf8.DecodeRune(word)

	fmt.Printf("%c %d\n", r, size)

	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Printf("%s = % [1]x\n", word)
}

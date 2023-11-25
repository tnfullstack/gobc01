// Indenting a text document
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	text := "Hi!, Go is awesome."
	text = Indent(text, 8)
	fmt.Println(text)

	text = Unindent(text, 3)
	fmt.Println(text)

	text = Unindent(text, 10)
	fmt.Println(text)

	text = IndentByRune(text, 10, '.')
	fmt.Println(text)
}

// Indent indenting the input by value and rune
func IndentByRune(input string, indent int, r rune) string {
	return strings.Repeat(string(r), indent) + input
}

// Indent indents text space by indent value
func Indent(input string, indent int) string {
	pad := indent + len(input)
	return fmt.Sprintf("% "+strconv.Itoa(pad)+"s", input)
}

// Unindent remove indentation space from text document
func Unindent(input string, indent int) string {
	count := 0
	for _, val := range input {
		if unicode.IsSpace(val) {
			count++
		}
		if count == indent || !unicode.IsSpace(val) {
			break
		}
	}
	return input[count:]
}

// embedding -
package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		text
		isbn string
	}

	moby := book{
		text: text{title: "moby dick", words: 206052},
		isbn: "1020305",
	}
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

}

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
		title string
		isbn  string
	}

	moby := book{
		text: text{title: "Moby Dick", words: 206052},
		isbn: "1020305",
	}
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

}

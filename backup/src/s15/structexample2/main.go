// structexample 2, compare struct value
package main

import "fmt"

// struct
type song struct {
	title  string
	artist string
}

func main() {
	s1 := song{
		title:  "Like a Rolling Stone",
		artist: "Bob Dylan",
	}

	s2 := song{
		title:  "What's going one",
		artist: "Marvin Gaye",
	}

	fmt.Printf("Song 1: %+v\nSong 2: %+v\n", s1, s2)

	// Compare the structs value
	if s1 == s2 {
		fmt.Println("equal!")
	}
}

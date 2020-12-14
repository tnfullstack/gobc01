package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `
English to Vietnamese dictionary

Usage: Enter command [enlish word] 
Example: dict hello
`

func main() {

	// Read command-line input
	args := os.Args[1:]

	// Varify input has two feilds (command + input)
	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	query := args[0]
	query = strings.ToLower(query)
	fmt.Println("User query =", query)

	// dictionary map that store [english] vietnamese key value pair
	english := map[string]string{
		"hello": "chào",
		"world": "thế giới",
		"up":    "lên",
		"down":  "xuống",
		"happy": "hạnh phúc",
		"laugh": "hạnh phúc",
	}

	// before delete map key/value
	fmt.Printf("Before delelte key value pair %q\n", english)

	// after delete map key/value entry
	delete(english, "laugh")
	fmt.Printf("After delelte key value pair %q\n", english)

	// create a map to hold vietnamese to english dict
	vietnamese := make(map[string]string)

	for k, v := range english {
		vietnamese[v] = k
	}

	fmt.Printf("What is in the english dict: %#v %v\n", english, len(english))
	fmt.Printf("What is in the vietnamese dict: %#v %v\n", vietnamese, len(vietnamese))

	if val, ok := english[query]; ok {
		fmt.Printf("%q means %q\n", query, val)
		return
	}

	if val, ok := vietnamese[query]; ok {
		fmt.Printf("%q means %q\n", query, val)
		return
	}
	fmt.Printf("Sorry, %q is not yet available in our dictionary!\n", query)
}

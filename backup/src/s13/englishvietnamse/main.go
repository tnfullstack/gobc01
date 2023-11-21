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
	dict := map[string]string{
		"hello": "chào",
		"world": "thế giới",
		"up":    "lên",
		"down":  "xuống",
		"happy": "hạnh phúc",
	}

	fmt.Printf("What is in the dictionary: %#v %v\n", dict, len(dict))

	// search for key value pair
	if val, ok := dict[query]; ok {
		fmt.Println("Query result:")
		fmt.Printf("%q means %q\n", query, val)
	} else {
		fmt.Printf("%s does not exist\n", query)
	}

	copied := map[string]string{"hello": "chào", "world": "thế giới"}

	first := fmt.Sprintf("%s", dict)
	second := fmt.Sprintf("%s", copied)

	if first == second {
		fmt.Printf("%s and %s are equal!\n", first, second)
	}
}

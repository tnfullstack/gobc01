// englishvietnamese - English to Vietnamese dictionary program
package main

import (
	"fmt"
	"os"
)

/*
// inefficient version using basic slices english and vietnamese to hold words
// this is O(n)
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: [English] to [Vietnamese]")
		return
	}

	input := args[0]

	english := []string{"good", "great", "perfect"}
	vietnamese := []string{"tốt", "tuyệt", "hoàn hoả"}

	for i, w := range english {
		if input == w {
			fmt.Printf("%q means %q\n", w, vietnamese[i])
			return
		}
	}
	fmt.Printf("%q not found\n", input)
}
*/

// Rewrite the program using maps

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: [English] to [Vietnamese]")
		return
	}

	input := args[0]

	dict := map[string]string{
		"good":    "tốt",
		"great":   "tuyệt",
		"perfect": "hoàn hảo",
		"tasty":   "ăn cơm", // this is a wrong translation
	}

	// we can add or replace existing word
	dict["up"] = "lên"
	dict["down"] = "xuống"
	dict["tasty"] = "ngon" // we can fix the wrong word by over write it
	dict["empty"] = ""

	// loop over the map
	for k, v := range dict {
		fmt.Printf("Dict data: %-5q means %q\n", k, v)
	}

	// Compare maps
	dictcopy := dict
	fmt.Println(dict, dictcopy) // these two looks they same but you can't campare them in current type

	// Convert the two to slice string
	first := fmt.Sprintf("%s", dict)
	second := fmt.Sprintf("%s", dictcopy)
	// Check the type and value
	fmt.Printf("first %T %q\n", first, first)
	fmt.Printf("second %T %q\n", second, second)
	fmt.Println(first == second)

	// Look up a word
	// Check for the existent of the word
	if result, ok := dict[input]; ok {
		fmt.Printf("%q means %q\n", input, result)
		return
	}
	fmt.Printf("%q not found\n", input)

}

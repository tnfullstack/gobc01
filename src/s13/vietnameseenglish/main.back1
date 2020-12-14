// vietnameseenglish - create Vietnamese to English dictionary program using existing englishvietnames
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: [Vietnamese] to [English]")
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
	dict["tasty"] = "ngon" // we can fix the wrong word by overwrite it
	dict["empty"] = ""

	// create a new map for vietnamese to english dictionary
	vietnamese := make(map[string]string, len(dict)) // len(dict) is not a required feature.

	// now clone the key and value pair from the dict by reverse the value to key and save
	// to vietnamese dictionary
	for k, v := range dict {
		vietnamese[v] = k
	}

	// Printout both dictionary to compare see outcome
	fmt.Println(dict)
	fmt.Println(vietnamese)

	if result, ok := vietnamese[input]; ok {
		fmt.Printf("%q means %q\n", input, result)
		return
	}
	fmt.Printf("%q not found\n", input)
}

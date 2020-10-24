package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

// Program usage information print when no command-line argument found.
const usage = `
Usage: command [house name]
here are the house's name: Gryffindor, Hufflepuf, Ravenclaw, Slytherin, and Bobo.
`

func main() {
	// read command-line arguments
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	// Create maps for House and Student Name
	house := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// search for student and print result
	search := strings.ToLower(args[0])
	// students := []string{}

	// Delete house if it not belong to Hogwarts
	if search == "bobo" {
		fmt.Printf("%s does not belong to Hogwarts, delete from list\n", search)
		deleteHouse(house, search)
		return
	}

	// If house exist and belong to Hogwarts print students in order
	if h, ok := house[search]; ok {
		sl := make([]string, len(h))

		// Copy the original slice house's student to temp list sl
		copy(sl, h)
		// Sort the temp sl (temp student list)
		sort.Strings(sl)

		fmt.Printf("~~~ %s Students ~~~\n", strings.Title(search))
		for _, n := range sl {
			// Print the sorted student list
			n = strings.Title(n)
			fmt.Printf("%-5s + %s\n", "", n)
		}
		// exit program after print
		return
	}
	// Print this if house not exist
	fmt.Printf("Sorry. I don't know anything about \"%s\"\n", search)

}

// Delete a house map[string][]string key
func deleteHouse(h map[string][]string, s string) {
	delete(h, s)
}

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

func main() {
	// House        Student Name
	// ---------------------------
	// gryffindor   weasley
	// gryffindor   hagrid
	// gryffindor   dumbledore
	// gryffindor   lupin
	// hufflepuf    wenlock
	// hufflepuf    scamander
	// hufflepuf    helga
	// hufflepuf    diggory
	// ravenclaw    flitwick
	// ravenclaw    bagnold
	// ravenclaw    wildsmith
	// ravenclaw    montmorency
	// slytherin    horace
	// slytherin    nigellus
	// slytherin    higgs
	// slytherin    scorpius
	// bobo         wizardry
	// bobo         unwanted

	// map for house data
	house := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// read command-line arguments
	agrs := os.Args[1:]

	// print usage info when there is no command argument
	if len(agrs) < 1 {
		fmt.Println("Enter command [house name]")
		fmt.Println("House Name: gryffindor, hufflepuf, ravenclaw, slytherin")
		return
	}

	// Save the search data in search variable
	search := agrs[0]

	// convert search string to lower case
	search = strings.ToLower(search)

	// Delete house bobo if there is a search for it.
	if search == "bobo" {
		delete(house, "bobo")
	}

	// search for key value in house map
	if name, ok := house[search]; ok {

		// house
		sort.Strings(name)
		fmt.Printf("~~~~~ %s student ~~~~~\n", search)
		for _, n := range name {
			fmt.Printf("%-10s %s\n", "", n)
		}
		return
	}
	fmt.Println("We don't know anything about", search)
}

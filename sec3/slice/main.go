// Slice
package main

import (
	"fmt"
	"sort"
)

func main() {

	// Slices data structure
	var animals []string

	// append assign string data to animals[0], animals[1]...
	animals = append(animals, "Dog")
	animals = append(animals, "Cat")
	animals = append(animals, "Bird")
	animals = append(animals, "Horse")

	fmt.Printf("animals type = %T, animals holds %s\n", animals, animals)

	// for loop cycle through the slice indexes
	fmt.Println("Using for loop to cycle through animals slice's indexes and printout data")
	for i, v := range animals {
		fmt.Printf("animals[%d] = %s\n", i, v)
	}

	// slicing allow selecting specific slice elements to display
	sliceI2 := animals[2:3] // sliceI2 will hold only element from animals[2]
	fmt.Printf("sliceI2 type = %T, sliceI2 = %s\n", sliceI2, sliceI2)

	sliceI012 := animals[0:3] // => animals[0, 1, 2]
	fmt.Printf("sliceI123 type = %T, sliceI123 = %s\n", sliceI012, sliceI012)

	last := animals[len(animals)-1]
	fmt.Printf("last type = %T, last = %s\n", last, last)

	// len return the slice number of elements
	elements := len(animals)
	fmt.Println("The length of animals =", elements)

	// sort will sort the element from the slice
	// isSorted check if the slice has been sorted and return bool value
	sort.Strings(animals)
	fmt.Println(animals, sort.StringsAreSorted(animals))
}

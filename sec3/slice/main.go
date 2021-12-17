// Slice
package main

import "fmt"

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
}

// Struct data type
package main

import "fmt"

type Car struct {
	Wheel int
	Door  int
	Speed float64
	Build string
	Model string
	Track bool
}

func main() {

	// Struct data type
	mustang := Car{
		Wheel: 4,
		Door:  2,
		Speed: 180.0,
		Build: "Ford",
		Model: "Mustang GT 5.0",
		Track: false,
	}
	fmt.Printf("%T, %v\n", mustang, mustang)

}

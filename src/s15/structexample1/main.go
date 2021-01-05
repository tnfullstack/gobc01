// Struct example 1
package main

import "fmt"

func main() {

	type person struct {
		fname, lname string
		age          int
	}

	mike := person{
		fname: "Mike",
		lname: "McNielon",
		age:   50,
	}

	tim := person{
		fname: "Tim",
		lname: "McLen",
		age:   59,
	}

	fmt.Println("Mike:", mike.fname, mike.lname, mike.age)
	fmt.Println("Tim:", tim.fname, tim.lname, tim.age)

	if mike.fname == tim.fname && mike.lname == tim.lname && mike.age == tim.age {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

// structbasic1 -
package main

import "fmt"

func main() {

	// declare person struct
	type person struct {
		firstName, lastName, job, education string
		age, income                         int
	}

	// declear two new persons chris and kim using person struct type
	var chris person
	var kim person

	// Initialize fields values
	chris.firstName = "Chris"
	chris.lastName = "Nguyen"
	kim.firstName = "Kim"
	kim.lastName = "Nguyen"
	chris.age = 50
	kim.age = 49

	// Declare a new person mike using struct literal
	mike := person{
		firstName: "Mike",
		lastName:  "Smith",
		age:       45,
		job:       "Tennis Coach",
		income:    98899,
		education: "College",
	}

	// Print the struct data to terminal
	fmt.Printf("chris: %#v\n", chris)
	fmt.Printf("kim: %#v\n", kim)
	fmt.Printf("Mike: %#v\n", mike)

	// Access idivitual field
	fmt.Printf("Mike's job: %v\n", mike.job)

}

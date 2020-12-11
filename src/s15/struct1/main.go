// structbasic1 -
package main

import "fmt"

// declare person struct
type person struct {
	firstName, lastName string
	job, education      string
	age, income         int
}

func main() {

	// declear two new persons chris and kim using person struct type
	chris := person{firstName: "Chris", lastName: "Nguyen", job: "", education: "", age: 50, income: 0}
	var kim person

	// Initialize fields values
	kim.firstName = "Kim"
	kim.lastName = "Nguyen"
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

// Expression
package main

import "fmt"

func main() {

	firstName := "Chris"
	lastName := "Ng"
	birthYear := 1975
	currYear := 2021
	job := "Programmer"
	hoby := "Tennis"
	var age int

	// Expression are firstName, lastName, job, and boby
	fmt.Printf("%s %s is a %s, his hoby is %s\n", firstName, lastName, job, hoby)

	// Expression are birthYear, currYear, firstName, and age from Printf line
	age = calcAge(birthYear, currYear)
	fmt.Printf("%s was born in %d and is now %d years old.\n", firstName, birthYear, age)

	// Expression is everything on the right side of the assignment operator :=
	isATeenager := age >= 13 && age <= 18 // Boolean operator
	fmt.Printf("is %s a teenager? => %t\n", firstName, isATeenager)

}

// function expression are by, cy, and age
func calcAge(by, cy int) (age int) {
	age = cy - by // expression
	return age
}

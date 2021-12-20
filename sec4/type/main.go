package main

import "fmt"

func main() {
	// Basic type (number, string, boolean)
	var num int

	// available types, but not widely use in production code
	// var num2 int16
	// var num3 int32
	// var num4 int64

	var num2 uint // can only hold possitive intergers 0, 1,2,3...

	// Floating point data types
	var num3 float32 // holding very large numbers on 32bits systems
	var num4 float64 // holding even larger numbers on 64bits systems

	num = 18
	num2 = 100
	num3 = 2999.99
	num4 = 88899.888

	// Display data value
	fmt.Println("These are some test values")
	fmt.Println("int", num)
	fmt.Println("uint", num2)
	fmt.Println("float32", num3)
	fmt.Println("float64", num4)

	// Strings data type, keep in mine string type are not immuable

	str1 := "" // string type with default empty string value
	fmt.Printf("Type: %T, value: %q\n", str1, str1)

	name := "Chris"
	fmt.Println("name variable hold", name)

	// Now assign/change Chris to Paul, keep in mind go created name = "Paul as a whole new string
	name = "Paul"
	fmt.Println("now name variable hold", name)
	fmt.Println("This is what you see, but Golang created a new variable to hold Paul.")

	// Boolean type
	var myBool bool

	fmt.Printf("Default bool type %T = %t\n", myBool, myBool)
	myBool = true
	fmt.Printf("Changed to true %T %t\n", myBool, myBool)

}

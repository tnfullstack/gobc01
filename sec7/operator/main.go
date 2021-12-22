// Operator
package main

import (
	"fmt"
	"math"
)

func main() {

	// math operator
	sum := 9 + 3*9
	fmt.Println("9 + 3 * 9 =", sum)

	sum = (9 + 3) * 9
	fmt.Println("(9 + 3) * 9 =", sum)

	// Area form mular a = pi * r * r
	var radius = 25

	// operator format 1
	area := math.Pi * float64(radius*radius)
	fmt.Println("if radius = 25, area =", area)

	// operator format 2
	r := float64(radius)
	area = math.Pi * r * r
	fmt.Println("if radius = 25, area =", area)

	// half
	half := 1 / 2 // output = 0
	fmt.Println("1 / 2 =", half)

	// operator format 1
	// convert to float
	result := 1.0 / 2.0
	fmt.Println("1 / 2 =", result)

	// math.Pow(3 3) => 3 to the power of 3 (math.Pow() always return a float)
	n1 := 3.0
	power := 3.0
	fmt.Println("3 to the power of 3 =", math.Pow(n1, power))
	power = 5
	fmt.Println("3 to the power of 5 =", math.Pow(n1, power))

	// Moduler operator %
	d := 5 / 3 // => 1 without remainder information
	fmt.Println("5 / 3 =", d)

	remainder := 5 % 3 // => 2 is the remainder when using the Moduler operator
	fmt.Println("5 % 3 =", remainder)

	// unary operator ++ or --
	a := 5
	a++
	fmt.Println("if a = 5, what is a++ in programming? answer a++ =", a)

	b := 5
	b--
	fmt.Println("b = 5, what is b--? answer b-- =", b)

}

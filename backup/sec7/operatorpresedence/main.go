// operators presedence
package main

import "fmt"

func main() {

	// Presedence
	// Multiplication and devision
	fmt.Println()
	fmt.Println("Multiplication and devision")
	a := 5.0 * 9.0 / 3.0
	b := (5.0 * 9.0) / 3.0
	c := 5.0 * (9.0 / 3.0)
	fmt.Println("a, b, c result =", a, b, c)

	// unclear
	fmt.Println()
	fmt.Println("unlear")
	unclear := 12.0 * (3 / 4)
	fmt.Println("unclear =", unclear)

	// parenthesis
	fmt.Println()
	fmt.Println("parenthesis")
	d := 12.0 / 3.0 / 4.0
	fmt.Println("d =", d)
	e := 12.0 / (3.0 / 4.0)
	fmt.Println("e =", e)

	// addition and subtraction
	f := 12 + 3 - 4
	g := (12 + 3) - 4
	h := 12 + (3 - 4)
	fmt.Println()
	fmt.Println("f, g, h =", f, g, h)

	i := 12 + 3*4
	j := (12 + 3) * 4
	k := 12 + (3 * 4)
	fmt.Println("i", i, "j", j, "k", k)
}

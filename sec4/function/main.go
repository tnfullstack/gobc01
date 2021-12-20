// Function, and function variatic
package main

import "fmt"

func main() {

	a := 15.0
	b := 9.0
	o1 := "+"
	o2 := "-"
	o3 := "*"
	o4 := "/"

	// Call a function
	fmt.Printf("%.2f %s %.2f = %.2f\n", a, o1, b, compute(a, o1, b))

	fmt.Printf("%.2f %s %.2f = %.2f\n", a, o2, b, compute(a, o2, b))

	fmt.Printf("%.2f %s %.2f = %.2f\n", a, o3, b, compute(a, o3, b))

	fmt.Printf("%.2f %s %.2f = %.2f\n", a, o4, b, compute(a, o4, b))

	sum := sumNums(1, 2, 3, 4, 5, 5, 6, 7, 8, 9)
	fmt.Println("Sum of 1,2,3,4,5,6,7,8,9 =", sum)

}

func compute(a float64, x string, b float64) (result float64) {

	switch x {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		fmt.Println("Sorry I can only compute basic (+, -, *, and /), please enter the right operation.")
	}

	return result
}

// Variatic function, sumNums will add a list of numbers, and return a total
func sumNums(n ...int) int {
	var sum int
	for _, x := range n {
		sum += x
	}

	return sum
}

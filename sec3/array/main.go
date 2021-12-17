// Aggregate data type
package main

import "fmt"

func main() {

	// Array data structure [5]string
	var animals [5]string
	animals[0] = "Dog"
	animals[1] = "Cat"
	animals[2] = "Bird"
	animals[3] = "Fish"
	animals[4] = "Monkey"
	fmt.Printf("%T %s, %s, %s, %s, %s\n", animals, animals[0], animals[1], animals[2], animals[3], animals[4])

	// Array [3]int, [2]float64, [2]byte
	num1 := [3]int{3, 5, 7}
	fmt.Printf("%T, %d\n", num1, num1)

	num2 := [2]float64{3.14, 8.55}
	fmt.Printf("%T, %.2f\n", num2, num2)

	b := [2]byte{'A', 'B'}
	fmt.Printf("%T, %q\n", b, b)

}

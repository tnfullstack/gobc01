// Pointers
package main

import "fmt"

func main() {

	// Pointers
	num := 288
	fmt.Println("num assign to 288 =", num)
	fmt.Println("num's memory address is", &num)

	var numPointer = &num
	fmt.Println("numPointer memory address is", &numPointer)
	fmt.Printf("numPointer type = %T, value = %v\n", numPointer, numPointer)

	*numPointer = 555

	fmt.Println("Assign 555 to *numPointer, will directy change the value in num, num =", num)
	fmt.Println("numPointer is", numPointer)

}

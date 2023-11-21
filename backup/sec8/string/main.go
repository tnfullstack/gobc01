// Working with strings
package main

import "fmt"

func main() {

	fmt.Println()
	str := "Thanh Nguyen"
	fmt.Println(str)
	fmt.Println()

	fmt.Println("Bytes")
	for i := range str {
		fmt.Printf("%d = %x\n", i, str[i]) // %x => hex number
	}

	fmt.Println()
	fmt.Println()
	fmt.Printf("%-10s %-10s %-10s\n", "Index", "Rune", "String")
	for j, s := range str {
		fmt.Printf("%-10d %-10d %-10c\n", j, s, s)
	}

	fmt.Println()
}

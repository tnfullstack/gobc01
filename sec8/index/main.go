// indexing
package main

import "fmt"

func main() {

	// Golang index count start from 0
	str := "Việt Nam bốn ngàn năm văn hiến"
	fmt.Println(str)
	fmt.Printf("The fist index str[0] holds %c\n", str[0]) // => V

	fmt.Printf("Length of str = %d, but index count start at 0 to 39\n", len(str))
	fmt.Printf("So first index 0 = %c, and last index 39 = %c\n", str[0], str[len(str)-1]) //
}

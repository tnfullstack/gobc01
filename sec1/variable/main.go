// Variable
package main

import "fmt"

func main() {
	var (
		str  = "Hello World!"
		str1 = "Go is the most fun and easy to learn language"
	)

	fmt.Println(str)

	saySomeThing(str1)
}

func saySomeThing(s string) {
	fmt.Println(s)
}

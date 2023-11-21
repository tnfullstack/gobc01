// Go program structure
package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher!")

	saySomeThing("Go is the most fun and simple to learn!")
}

func saySomeThing(s string) {
	fmt.Println(s)
}

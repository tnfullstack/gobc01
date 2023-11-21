// Go routin, and channel
package main

import "fmt"

func main() {

	go doSomeThing("Hello World!")

	fmt.Println("This is from main.")

	for {
		// fmt.Println(i, "Test")
	}

}

func doSomeThing(s string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("S is", s)
	}
}

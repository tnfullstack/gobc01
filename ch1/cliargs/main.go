package main

import (
	"fmt"
	"os"
)

func main() {
	// Command line argument
	args := os.Args

	// This call will print
	// all command line arguments.
	fmt.Println(args)

	// The first arugment, zero item from slice,
	// is the name of the called binary.
	argName := args[0]
	fmt.Printf("The binary name is: %s\n", argName)

	// The rest of the arguments could be obtained
	// by omitting the first argument.
	otherArgs := args[1:]
	fmt.Println("Other arguments:", otherArgs)

	for i, a := range args {
		fmt.Printf("CLI arguments: %d, %s\n", i, a)
	}
}

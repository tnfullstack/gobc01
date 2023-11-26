// Opening a file by name
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please enter the file name")
		return
	}

	file := args[0]

	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("failed to open file %s\n", err)
	}
	defer f.Close()

	c, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("### file content ###\n%s\n", string(c))

	f, err = os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, "Test String")
	f.Close()

}

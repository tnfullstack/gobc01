// Reading the file into a string
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("### Read as reader ###")
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("### ReadFile ###")
	// for smaller files
	file, err := os.ReadFile("temp/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Last operation", string(file))
}

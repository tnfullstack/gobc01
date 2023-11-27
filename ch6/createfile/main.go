// Creating and writing to files
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString("This text is going into the newly created file named test.txt! ")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, strings.NewReader("Additional text going into text.txt file."))
	if err != nil {
		log.Fatal(err)
	}
}

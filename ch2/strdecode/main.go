// Decoding a string from the non-Unicode charset
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	// Open windows-1250 file
	file, err := os.Open("win1250.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read all in raw form.
	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	content := string(b)

	fmt.Println("Without decode: " + content)

	// Decode to unicode
	decoder := charmap.Windows1250.NewDecoder()
	reader := decoder.Reader(strings.NewReader(content))
	b, err = io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deocded: " + string(b))
}

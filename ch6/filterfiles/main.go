// Filtering file listings
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	for i := 1; i <= 6; i++ {
		f, err := os.Create(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", f.Name())
	}

	m, err := filepath.Glob("./test.file[1-3]")
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range m {
		fmt.Println(val)
	}

	// Cleanup
	for i := 1; i <= 6; i++ {
		err := os.Remove(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}
}

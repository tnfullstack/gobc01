package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// Path to executable file
	fmt.Println("Executable file:", ex)

	// Resolve the directory of the executable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path:" + exPath)

	// Use EvalSymlinks to get the real path
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Symlink evaluated:" + realPath)
}

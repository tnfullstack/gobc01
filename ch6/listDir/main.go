// Listing a directory
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: [command] [dir path] [Enter]")
		return
	}
	input := args[0]

	fmt.Println("List by ReadDir")
	listDirByReadDir(input)
	fmt.Println()
	fmt.Println("List by Walk")
	listDirByWalk(input)
}

func listDirByReadDir(path string) {
	lst, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}

func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		// Walk the given dir without printing out.
		if wPath == path {
			return nil
		}

		// If given path is folder stop list recursively and print as folder.
		if info.IsDir() {
			fmt.Printf("[%s]\n", wPath)
			return filepath.SkipDir
		}

		// Print file name
		if wPath != path {
			fmt.Println(wPath)
		}
		return nil
	})
}

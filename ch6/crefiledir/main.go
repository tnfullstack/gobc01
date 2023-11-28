// Creating files and directories
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("created.file")
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	f, err = os.OpenFile("created.file", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	err = os.Mkdir("createDir", 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll("sampleDir/path1/path2", 0777)
	if err != nil {
		log.Fatal(err)
	}
}

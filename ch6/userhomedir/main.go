// Resolving the user home directory
package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User home directory: %s\n", usr.HomeDir)
}

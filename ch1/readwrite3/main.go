// Reading/Writing from the child process
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		if sc.Text() == "q" {
			os.Exit(-1)
		}
		fmt.Println(sc.Text())
	}
}

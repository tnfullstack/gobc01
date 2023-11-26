// Reading command-line input using bufio.NewScanner and os.Stdin
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("--> ")
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
		fmt.Print("--> ")
		if strings.ToLower(txt) == "q" {
			os.Exit(-1)
		}
	}
}

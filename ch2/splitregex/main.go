// Splitting string with the help of Regular Expression (regex)
package main

import (
	"fmt"
	"regexp"
)

func main() {
	const testStr = "Mary*had,a%little_lam"
	const testStr2 = "Dang!Cong&San^Viet@Nam#Muon$Nam."

	regexpCompiled := regexp.MustCompile("[*,%_!&@#$^()]{1}")

	list1 := regexpCompiled.Split(testStr, -1)
	list2 := regexpCompiled.Split(testStr2, -1)

	for i, w := range list1 {
		fmt.Printf("Word %d: is %s\n", i, w)
	}

	for i, w := range list2 {
		fmt.Printf("Word %d: is %s\n", i, w)
	}
}

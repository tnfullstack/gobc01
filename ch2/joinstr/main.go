// Join the String slice with separator
package main

import (
	"fmt"
	"strings"
)

func main() {
	const selectBase = "SELECT * FROM use WHERE %s "

	testStr := []string{" FIRST_NAME = 'Jack' ",
		" INSURANCE_NO = 33322333 ",
		" EFFECTIVE_FROM = SYSDATE ",
	}

	joinStr := strings.Join(testStr, "AND")
	fmt.Printf(selectBase+"\n", joinStr)
}

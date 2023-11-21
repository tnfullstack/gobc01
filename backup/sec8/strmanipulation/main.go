// String manipulation
package main

import (
	"fmt"
	"strings"
)

func main() {

	str := []string{
		"Việt Nam bốn ngày năm văn hiến.",
		"Sài Gòn là thủ đô của miền nam Việt Nam.",
		"Huế là thủ đô của miền trung.",
		"Hà Nội là thủ đô của miền bắc.",
	}

	// var newstr []string
	for i, w := range str {
		// strings.Replace
		str[i] = strings.Replace(w, "miền", "phương", -1)
		fmt.Println(w, "--- Before (replace miền to phương)")
		fmt.Println(str[i], "--- After")
	}

	fmt.Println()
	// string comparison
	answer := "ABC" > "ABD"
	fmt.Println("ABC > ABD", answer)

	fmt.Println()
	// strings.TrimSpace()
	email := " chris@gmail.com   "
	fmt.Printf("%q\n", email)
	email = strings.TrimSpace(email)
	fmt.Printf("%q\n", email)

}

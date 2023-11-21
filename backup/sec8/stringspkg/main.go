// using strings package
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

	// search for the the word Việt
	count := 0
	// var newstr []string
	for i, w := range str {
		if strings.Contains(w, "Việt") {
			fmt.Printf("We found %s in string %q, index #%d\n", "Việt", w, strings.Index(w, "Việt"))
		}
		if strings.HasPrefix(w, "Sài") {
			fmt.Printf("%s found in %q\n", "Sài", w)
		}
		if strings.HasSuffix(w, "bắc.") {
			fmt.Printf("Found %s in %q\n", "bắc.", w)
		}
		count += strings.Count(w, "thủ")

		// strings.Replace
		str[i] = strings.Replace(w, "miền", "phương", -1)

	}

	// Print total count
	fmt.Printf("Found %d count of thủ\n", count)

	fmt.Println()
	for _, w := range str {
		fmt.Println(w)
	}

}

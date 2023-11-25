// Finding the substring in text using regex pattern
package main

import (
	"fmt"
	"regexp"
)

func main() {
	const testStr = `[{ \"email\": \"email@domain.com\" \"phone\": 5554567890},
                      { \"email\": \"other@domain.com\" \"phone\": 5554567890}]`

	// This pattern is simplified for brevity
	emailRegexp := regexp.MustCompile(`[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\.[a-z]{1,}`)

	first := emailRegexp.FindString(testStr)
	fmt.Println("First:", first)

	all := emailRegexp.FindAllString(testStr, -1)
	fmt.Println("All: ")
	for _, val := range all {
		fmt.Println(val)
	}
}

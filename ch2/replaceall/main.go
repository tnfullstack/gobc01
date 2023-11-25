// Replacing all of the string with testStr.ReplaceAllString(s string, "Replacement string")
package main

import (
	"fmt"
	"regexp"
)

func main() {
	testStr := "Mary has a little lam"

	testStrSlice := []string{"adam[23]", "eve[7]", "Job[48]", "job[48]", "snakey", "snakey[1234567890]"}

	regexPatten := `^[a-z]+\[[0-9]+\]$`

	regex := regexp.MustCompile("l[a-z]+")
	out := regex.ReplaceAllString(testStr, "Replace all string with new string.")
	fmt.Println(out)

	fmt.Println(testStr, matchTest(regexPatten, testStr))

	for i, str := range testStrSlice {
		fmt.Println(i, str, matchTest(regexPatten, str))
	}
}

func matchTest(regexStr string, matchStr string) bool {
	regCompiled := regexp.MustCompile(regexStr)
	return regCompiled.MatchString(matchStr)
}

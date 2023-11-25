// Managing whitespace in a string
package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	testStr := "\t\t\n Go \tis\t Awesome \t\t"
	resultStr := strings.TrimSpace(testStr)
	fmt.Println(resultStr)

	strWithWhiteSpace := "\t\t\n Go \tis\n Awesome \t\t"
	regexCompiled := regexp.MustCompile(`\s+`)
	resultStr = regexCompiled.ReplaceAllString(strWithWhiteSpace, " ")
	fmt.Println(resultStr)

	testStr2 := "Need space here"
	fmt.Println(pad(testStr2, 15, "CENTER"))
	fmt.Println(pad(testStr2, 15, "LEFT"))

}

func pad(input string, padLen int, align string) string {
	inputLen := len(input)

	if inputLen >= padLen {
		return input
	}

	repeat := padLen - inputLen
	var output string

	switch align {
	case "RIGHT":
		output = fmt.Sprintf("%s"+strconv.Itoa(-padLen)+"s", input) // Not sure how this works
	case "LEFT":
		output = fmt.Sprintf("% "+strconv.Itoa(padLen)+"%s", input)
	case "CENTER":
		bothRepeat := float64(repeat) / float64(2)
		left := int(math.Floor(bothRepeat)) + inputLen
		right := int(math.Ceil(bothRepeat))
		output = fmt.Sprintf("% "+strconv.Itoa(left)+"s%"+strconv.Itoa(right)+"s", input, "")
	}
	return output
}

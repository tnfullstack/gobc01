// String Capitalization manipulation
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	const email = "ExamPle@domain.Com"
	const name = "isaac newton"
	const upc = "upc"
	const i = "i"

	const snakeCase = "fist_name"

	// For comparing the use input sometime it is better to compare
	// the input in the same case.
	input := "Example@domain.cOm"
	input = strings.ToLower(input)

	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("Email matches: %t\n", matches)

	upcCode := strings.ToUpper(upc)
	fmt.Println("UPPER case: " + upcCode)

	// This digraph has different upper case and tittle case.
	str := "dz"
	fmt.Printf("%s in upper: %s and title: %s \n", str, strings.ToUpper(str), strings.ToTitle(str))

	// Use of XXXSpecial function
	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("ToTitle is defferent: %#U vs. %#U \n", title[0], []rune(titleTurk)[0])
	}

	// In some cases the input needs to be corrected in case.
	// correctNameCase := strings.Title(name) -> deprecated
	correctNameCase := strings.ToTitle(name)
	fmt.Println("Corrected name: " + correctNameCase)

	// Converting the snake case to camel case with the use of Title and ToLower functions.
	firstNameCamel := toCamelCase(snakeCase)
	fmt.Println("Camel case: " + firstNameCamel)

}

func toCamelCase(testStr string) string {
	splitStr := strings.Fields(strings.ToLower(strings.Replace(testStr, "_", " ", -1)))
	var finalStr string
	for _, str := range splitStr {
		finalStr += strings.ToTitle(str[:1]) + str[1:]
	}
	return finalStr
}

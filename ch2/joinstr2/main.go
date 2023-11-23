// Join string slice with separator
package main

import (
	"fmt"
	"strings"
)

type JoinFunc func(piece string) string

func JoinWithFunc(strSlice []string, jF JoinFunc) string {
	concatenate := strSlice[0]
	for _, v := range strSlice[1:] {
		concatenate += jF(v) + v
	}
	return concatenate
}

func main() {
	const selectBase = "SELECT * FROM use WHERE "

	testStr := []string{
		" FIRST_NAME = 'Jack' ",
		" INSURANCE_NO = 33322333 ",
		" EFFECTIVE_FROM = SYSDATE ",
	}

	joinFunc := func(p string) string {
		if strings.Contains(p, "INSURANCE") {
			return "OR"
		}
		return "AND"
	}

	result := JoinWithFunc(testStr, joinFunc)
	fmt.Println(selectBase + result)
}

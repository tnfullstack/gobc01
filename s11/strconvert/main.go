// string convert, index, slice, bytes, runes and strings
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Thành Phố 🥎"

	bytes := []byte(str)
	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	fmt.Println()
	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	fmt.Println()
	fmt.Printf("1st byte   : %c\n", str[0])
	fmt.Printf("2nd byte   : %c\n", str[1])
	fmt.Printf("3rd rune   : %s\n", str[2:4])
	fmt.Printf("4th byte   : %c\n", str[4])
	fmt.Printf("5th byte   : %c\n", str[5])
	fmt.Printf("6th byte   : %c\n", str[6])
	fmt.Printf("7th byte   : %c\n", str[7])
	fmt.Printf("8th byte   : %c\n", str[8])
	fmt.Printf("9th rune   : %s\n", str[9:13])
	fmt.Printf("12th byte  : %c\n", str[14])
	fmt.Printf("13th rune  : %s\n", str[len(str)-4:])

	runes := []rune(str)
	fmt.Println()
	fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d runes\n", len(runes))
	fmt.Printf("1st rune    : %c\n", runes[0])
	fmt.Printf("2nd rune    : %c\n", runes[1])
	fmt.Printf("first five  : %c\n", runes[:5])
	fmt.Printf("All runes   : %c\n", runes[:])
}

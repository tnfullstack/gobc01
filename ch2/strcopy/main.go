// Concatenating string using copy
package main

import "fmt"

func main() {
	str := []string{"This ", "is ", "even ", "more ", "performant "}

	bs := make([]byte, 1*1024)
	bl := 0

	for _, val := range str {
		bl += copy(bs[bl:], []byte(val))
	}
	fmt.Println(string(bs))
	fmt.Printf("%s length %d bytes\n", "bs", len(bs))
	fmt.Printf("%s capacity %d bytes\n", "bs", cap(bs))
	fmt.Printf("bs bytes used: %d, bytes available %d\n", bl, len(bs)-bl)
}

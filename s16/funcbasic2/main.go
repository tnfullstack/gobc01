package main

import (
	"fmt"
	"strconv"
)

func main() {
	local := 10

	a := 2
	b := 20

	local = show(local)
	fmt.Println(local)

	p := multiplication(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, p)

	r, err := incrByStr(b, "50")
	if err != nil {
		fmt.Println("Wrong input value.", err)
		return
	}
	fmt.Printf("%d and %s = %d\n", b, "50", r)
}

func show(n int) int {
	n += 10
	fmt.Println(n)

	return n
}

func multiplication(a, b int) int {
	return a * b
}

func incrByStr(n int, s string) (int, error) {
	m, err := strconv.Atoi(s)

	n = multiplication(n, m)
	return n, err

}

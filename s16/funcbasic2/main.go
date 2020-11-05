package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Declare a local variable
	local := 10
	// call/send local = 10 to show() function
	show(local) // Show n = 10

	// call/pass 10 to incrWrong()
	incrWrong(local)
	fmt.Printf("main -> local = %d\n", local) // local variable is still = 10

	// call/pass 10 to addOne()
	local = addOne(local)
	show(local)

	// call/pass 10 to multiplyByFive
	local = multiply(local, 5)
	show(local)

	// call/pass local and string "90" to incrByStr()
	local, err := incrByStr(local, "90")
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	show(local)

	//
	local = sanitize(incrByStr(local, "90"))
	show(local)

	// call/pass local and limit value to limit()
	local = limit(local, 50)
	show(local)
}

// Print any integer value passed in.
func show(n int) {
	fmt.Printf("Show -> local = %d\n", n)
}

// crease n by one n++ = n + 1, but value only stay inside function scope
func incrWrong(n int) {
	n++
	// n++ is local to incrWrong() -> Print n will display n + 1
	// fmt.Println("n from incrWrong = ", n)
}

// increase n by one and return value
func addOne(n int) int {
	n++
	return n
}

// mulpliply n by factor of 7 and return result
func multiply(n, f int) int {
	return n * f
}

// convert a string value to integer
func incrByStr(a int, s string) (int, error) {
	b, err := strconv.Atoi(s)
	a = multiply(a, b)
	return a, err
}

// return n or 0 if there is error value
func sanitize(n int, err error) int {
	if err != nil {
		return 0
	}
	return n
}

// return the limit value
func limit(n, lim int) (m int) {
	// var m int
	m = n
	if m >= lim {
		m = lim
	}
	return
}

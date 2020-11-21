package main

import "fmt"

func show() {
	if N == 0 {
		return
	}

	fmt.Printf("Show N : is %d\n", N)
}

func incrN() {
	N++
}

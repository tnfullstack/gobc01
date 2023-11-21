package main

import "fmt"

func showN() {
	if N == 0 {
		return
	}
	fmt.Printf("Show N	: N is %d\n", N)
}

func incrN() {
	N++
}

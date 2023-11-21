// logparserv2
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {

		//
		p.lines++

		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		// passing data to update function
		p = update(p, parsed)

	}

	// Printing domains and visits
	fmt.Printf("%-25s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 40))

	// sort slice domains[]string
	sort.Strings(p.domains)

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-25s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-25s %10d\n", "TOTAL", p.total)

	// handle the error
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}

}

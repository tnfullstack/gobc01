package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {

	var (
		domains []string
		lines   int
		total   int
	)

	// read data from bufio.Newscanner(os.Stdin)
	in := bufio.NewScanner(os.Stdin)

	domainvisit := make(map[string]int)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())

		domain := fields[0]
		if len(fields) < 2 {
			fmt.Printf("Wrong data %v line # %d\n", fields, lines)
			return
		}

		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("Wong logs data %d %d\n", visits, lines)
			return
		}

		if _, ok := domainvisit[domain]; ok {
			domains = append(domains, domain)
		}

		domainvisit[domain] += visits

		total += visits
	}

	// Header
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISIT")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)

	for _, d := range domains {
		v := domainvisit[d]
		fmt.Printf("%-30s %10d\n", d, v)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)
}

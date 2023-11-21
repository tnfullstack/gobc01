// logparser -
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// declare variable for map sum = map[string]ing and slice domains[]string
	var (
		sum     map[string]int
		domains []string
		total   int
		lines   int
	)

	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++
		fields := strings.Fields(in.Text())
		if len(fields) < 2 {
			fmt.Printf("Wrong input: %v (line #%d)\n", fields, lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("Wrong input: %v (line #%d)\n", fields, lines)
			return
		}

		// look for domain in map sum[domain], if domain does not exist, append to slice domains
		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		// assign number of visits to map sum[domain]: visists
		sum[domain] += visits

		// total visist for all domains
		total += visits
	}

	// Printing domains and visits
	fmt.Printf("%-25s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 40))

	// sort slice domains[]string
	sort.Strings(domains)

	for _, d := range domains {
		v := sum[d]
		fmt.Printf("%-25s %10d\n", d, v)
	}
	fmt.Printf("\n%-25s %10d\n", "TOTAL", total)

}

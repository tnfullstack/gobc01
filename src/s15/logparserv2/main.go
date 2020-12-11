// logparserv2
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func main() {

	p := parser{
		sum: make(map[string]result),
	}

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		fields := strings.Fields(in.Text())

		if len(fields) < 2 {
			fmt.Printf("Wrong input %v (line #%d)\n", fields, p.lines)
			return
		}

		domain := fields[0]

		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("Wrong input %d (line #%d)\n", visits, p.lines)
			return
		}

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		// Total visits from all websites
		p.total += visits
		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}

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

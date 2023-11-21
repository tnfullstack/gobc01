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
	sum     map[string]result // total visits per domain
	domains []string          // unique domain names
	total   int               // total to all domains
	lines   int               // number of parsed lines
}

func main() {

	p := parser{
		sum: make(map[string]result),
	}

	// read data from a file with standard input (command < fileName.txt)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		fields := strings.Fields(in.Text())

		if len(fields) < 2 {
			fmt.Printf("Wrong input %v (line #%d)\n", fields, p.lines)
			return
		}

		// store domain name from fields[0] to domain variable
		d := fields[0]

		if _, ok := p.sum[d]; !ok {
			p.domains = append(p.domains, d)
		}

		v, err := strconv.Atoi(fields[1])
		if v < 0 || err != nil {
			fmt.Printf("Wrong input %d (line #%d)\n", v, p.lines)
			return
		}
		// fmt.Printf("%T %s %T %d\n", d, d, v, v)

		// Store domain name and number of visit per website to parser {sum: map[string]}
		/*
			r := result{
				domain: d,
				visits: v + p.sum[d].visits,
			}
		*/

		p.sum[d] = result{
			domain: d,
			visits: v + p.sum[d].visits,
		}

		p.total += v

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

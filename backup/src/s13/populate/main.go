package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Populate and Lookup
//
//  Add elements to the maps that you've declared in the
//  first exercise, and try them by looking up for the keys.
//
//  Either use the `make()` or `map literals`.
//
//  After completing the exercise, remove the data and check
//  that your program still works.
//
//
//  1. Phone numbers by last name
//     --------------------------
//     bowen  202-555-0179
//     dulin  03.37.77.63.06
//     greco  03489940240
//
//     Print the dulin's phone number.
//
//
//  2. Product availability by Product ID
//     ----------------
//     617841573 true
//     879401371 false
//     576872813 true
//
//     Is Product ID 879401371 available?
//
//
//  3. Multiple phone numbers by last name
//     ------------------------------------------------------
//     bowen  [202-555-0179]
//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
//     greco  [03489940240 03489900120]
//
//     What is Greco's second phone number?
//
//
//  4. Shopping basket by Customer ID
//     -------------------------------
//     100 [617841573:4 576872813:2]
//     101 [576872813:5 657473833:20]
//     102 []
//
//     How many of 576872813 the customer 101 is going to buy?
//                (Product ID)  (Customer ID)
//
//
// EXPECTED OUTPUT
//
//   1. Run the solution to see the output
//   2. Here is the output with empty maps:
//
//      dulin's phone number: N/A
//      Product ID #879401371 is not available
//      greco's 2nd phone number: N/A
//      Customer #101 is going to buy 5 from Product ID #576872813.
//
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Example: command [value]")
		return
	}

	query := args[0]
	query = strings.ToLower(query)

	// 1. Phone numbers by last name
	phone := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	if val, ok := phone[query]; ok {
		fmt.Printf("%q phone #: %s\n", query, val)
		return
	}

	// 2. Product availability by Product ID
	product := map[int]bool{
		1001: true,
		1002: false,
		1003: true,
	}

	// 3. Multiple phone numbers by last name
	multiPhone := map[string][]string{
		"smith":    {"123-111-1111", "555-555-5555"},
		"jones":    {"123-222-3333"},
		"williams": {"123-333-4444"},
		"taylor":   {"123-444-5555"},
		"davis":    {"123-555-6666"},
		"brown":    {"123-666-7777"},
	}
	fmt.Printf("%v\n", multiPhone)

	// 4. Shopping basket by Customer ID
	brasket := map[int]map[int]int{
		101: {1001: 3, 1002: 2, 1003: 5},
		102: {1005: 5, 1006: 4},
		103: {1007: 5, 1008: 6, 1009: 7},
	}
	fmt.Printf("%v\n", brasket)

	pID, err := strconv.Atoi(query)

	if err != nil {
		fmt.Println("Cannot find item")
		return
	}

	if val, ok := product[pID]; ok {
		if val == true {
			fmt.Printf("Product %d is in tock.\n", pID)
			return
		}
	}
	fmt.Printf("Sorry, I can not find %d.\n", pID)

}

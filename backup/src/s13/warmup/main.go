// package main
package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {

	/*
		args := os.Args[1:]

		if len(args) < 1 {
			fmt.Println("Example: command [value]")
			return
		}

		query := args[0]
		query = strings.ToLower(query)
	*/

	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number

	phone := map[string]string{
		"smith":    "123-111-1111",
		"jones":    "123-222-3333",
		"williams": "123-333-4444",
		"taylor":   "123-444-5555",
		"davis":    "123-555-6666",
		"brown":    "123-666-7777",
	}
	fmt.Printf("%v\n", phone)

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	product := map[int]bool{
		001: true,
		002: true,
		003: true,
		004: true,
		005: false,
	}
	fmt.Printf("%v\n", product)

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	phone1 := map[string][]string{
		"smith":    {"123-111-1111", "555-555-5555"},
		"jones":    {"123-222-3333"},
		"williams": {"123-333-4444"},
		"taylor":   {"123-444-5555"},
		"davis":    {"123-555-6666"},
		"brown":    {"123-666-7777"},
	}
	fmt.Printf("%v\n", phone1)

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	// val, _ := phone[query]
	brasket := map[string]map[string]int{
		"Mis. X":   {"Banana": 3, "Icescream": 2, "Egg": 5},
		"Mr. Mike": {"Meat Ball": 5, "Apple": 4},
		"Mr. Tim":  {"Orange": 5, "Mango": 6, "Cherry": 7},
	}
	fmt.Printf("%v\n", brasket)
}

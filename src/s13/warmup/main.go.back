// warmup - warmup exercise
package main

import "fmt"

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
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	phonebook := map[string]string{
		"smith":    "123-111-1111, 123-777-8888",
		"jones":    "123-222-3333",
		"williams": "123-333-4444",
		"taylor":   "123-444-5555",
		"davis":    "123-555-6666",
		"brown":    "123-666-7777, 213-999-0001",
	}
	phonebook["Nguyen"] = "510-566-9795"

	// Print map phonebook
	fmt.Printf("%q\n", phonebook)

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	products := map[int]bool{
		0001: true,
		0002: false,
	}
	products[0003] = true
	fmt.Printf("%v\n", products)
	fmt.Printf("Product ID: 0001 is in stock? %t\n", products[0001])

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	multiphones := map[string][]string{
		"smith": {
			"222-333-4444",
			"222-333-4401",
			"222-333-4401",
		},
		"jone": {
			"222-333-5501",
			"222-333-5502",
		},
	}
	fmt.Printf("%v\n", multiphones)

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	customerID := map[string]map[string]int{
		"C0001": map[string]int{
			"P0001": 1,
			"P0002": 2,
			"P0003": 5,
			"P0004": 2,
			"P0005": 3,
		},
		"C0002": map[string]int{
			"P0001": 3,
			"P0002": 4,
		},
	}
	customerID["C0003"] = map[string]int{}
	fmt.Printf("%v\n", customerID)

}

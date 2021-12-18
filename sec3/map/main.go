// Map data structure
package main

import "fmt"

func main() {
	// Map data structure
	var months = make(map[string]int)

	months = map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Arp": 4,
		"May": 5,
		"Jun": 6,
		"Jul": 7,
		"Aug": 8,
		"Sep": 9,
		"Oct": 10,
		"Nov": 11,
		"Dec": 12,
	}
	fmt.Printf("months type is %T, value = %d, %d\n", months, months["Jan"], months["Feb"])

	// Adding element to map
	months["Test"] = 13

	// Check for map element
	element, ok := months["Test"]
	if ok {
		fmt.Println("element found", element)
	} else {
		fmt.Println(element, "element not exist")
	}

	// Delete element from map
	delete(months, "Test")

	// Check for map element
	element, ok = months["Test"]
	if ok {
		fmt.Println("element found", element)
	} else {
		fmt.Println(element, "element not exist")
	}

	// Over write map element
	months["Jan"] = 202201
	fmt.Println("Jan =", months["Jan"])

	// Iterate through map
	for k, v := range months {
		fmt.Printf("Key = %s, value = %d\n", k, v)
	}
}

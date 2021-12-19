// Boolean Expression operator
package main

import "fmt"

func main() {
	lbToOz := 16 // 1lb = 16oz
	gold := 5    // 5 pounds
	copper := 5  // 5 pounds

	goldPrice := 1798.60 // 1798.60/ounce base on Dec 18, 2021 price
	copperPrice := .27   // 0.27/ounce base on Dec 18, 2021 price

	fmt.Printf("%d pounds of gold is heavier than %d pounds of copper => %t\n", gold, copper, gold > copper)
	fmt.Printf("%d pounds of gold is lighter than %d pounds of copper => %t\n", gold, copper, gold < copper)
	fmt.Printf("%d pounds of gold is heavier or equal to %d pounds of copper => %t\n", gold, copper, gold <= copper)
	fmt.Printf("%d pounds of gold is equal to %d pounds of copper => %t\n", gold, copper, gold == copper)
	fmt.Println()

	// Convert pound to ounce 1lb = 16oz
	goldOz := gold * lbToOz
	copperOz := copper * lbToOz

	// Calculate total coast for gold and copper and compare in dollars amount
	goldCost := float64(goldOz) * goldPrice
	copperCost := float64(copperOz) * copperPrice

	// Display total cost for gold and copper
	fmt.Printf("%d lbs of gold cost $%.2f\n", gold, goldCost)
	fmt.Printf("%d lbs of copper cost $%.2f\n", copper, copperCost)
	fmt.Println()
	fmt.Println("Gold is less expensive than copper =>", goldCost < copperCost)
	fmt.Println("Gold is more expensive than copper =>", goldCost > copperCost)
	fmt.Println("Gold cost the same as copper =>", goldCost <= copperCost)
	fmt.Println("Gold cost more or equal to copper =>", goldCost >= copperCost)

}

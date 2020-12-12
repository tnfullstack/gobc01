package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

type item struct {
	name      string
	id, price int
}

type game struct {
	item
	genre string
}

func main() {

	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{
			item:  item{id: 2, name: "x-com", price: 30},
			genre: "strategy",
		},
		{
			item:  item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}

	fmt.Printf("%-5s %-15s %-10s %-20s\n", "ID", "Name", "Price", "Genre")
	fmt.Println(strings.Repeat("-", 50))

	for _, g := range games {
		fmt.Printf("%-5d %-15s %-10d %-20s\n", g.id, g.name, g.price, g.genre)

	}
}

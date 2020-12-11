package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
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

	fmt.Printf("This game store has %d games. %#v\n", len(games), games)
	// fmt.Printf("%d %d\n", games.item[1], games.item[2])

	// id := games[0].id

	fmt.Println("Test", games[0])

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
> list : list all games
> quit : exit program
`)

		if !in.Scan() {
			break
		}

		switch in.Text() {
		case "quit":
			fmt.Println("Bye!")
			return
		case "list":
			fmt.Printf("%-5s %-15s %-10s %-20s\n", "ID", "Name", "Price", "Genre"+"\n")
			for _, g := range games {
				fmt.Printf("%-5d %-15s %-10d %-20s\n", g.id, g.name, g.price, g.genre)
			}
		}
	}
}

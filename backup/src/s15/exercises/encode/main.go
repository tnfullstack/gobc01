// encode
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

type item struct {
	Name      string
	ID, Price int
}

type game struct {
	item
	Genre string
}

func main() {

	// Create to new map
	games := []game{
		{
			item: item{
				ID:    1,
				Name:  "god of war",
				Price: 50,
			},
			Genre: "action adventure",
		},
		{
			item:  item{ID: 2, Name: "x-com", Price: 30},
			Genre: "strategy",
		},
		{
			item:  item{ID: 3, Name: "minecraft", Price: 20},
			Genre: "sandbox",
		},
	}

	searchID := make(map[int]game)
	for _, g := range games {
		searchID[g.ID] = g
	}

	fmt.Printf("This game store has %d games.\n", len(games))

	for {
		fmt.Printf(`
> list : list all games
> id   : list by id, example "id 2"
> save : save games list to json
> quit : exit program
`)

		in := bufio.NewScanner(os.Stdin)
		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		// fmt.Printf("%T, %v\n", cmd, len(cmd))

		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit", "q":
			fmt.Println("Program exited, Bye!")
			return

		case "list", "l":
			fmt.Printf("%-5s %-15s %-10s %s\n", "ID", "Name", "Price", "Genre"+"\n")
			for _, g := range games {
				fmt.Printf("%-5d %-15s $%-10d %s\n", g.ID, g.Name, g.Price, g.Genre)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("Wrong id")
				continue
			}

			num, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("ID should be a number.")
				continue
			}

			g, ok := searchID[num]
			if !ok {
				fmt.Println("Sorry, the game is not available!")
				continue
			}
			fmt.Printf("#%d: %-15q %-20s $%d\n", g.ID, g.Name, "("+g.Genre+")", g.Price)
		case "save", "s":
			out, err := json.MarshalIndent(games, "", "\t")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(string(out))
			return
		}
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]
`

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {

	// encodable and decodable game type
	type jsonGame struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
	}

	// load the initial data from json
	var decoded []jsonGame

	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem:", err)
		return
	}

	// fmt.Println(decoded)

	var games []game
	for _, dg := range decoded {
		games = append(games, game{item{dg.ID, dg.Name, dg.Price}, dg.Genre})
	}

	searchID := make(map[int]game)

	for _, g := range games {
		searchID[g.id] = g
	}

	// Print the game store announcement
	fmt.Printf("\nThis game store has %d games\n", len(games))

	for {
		fmt.Println(`
> list : list all games
> id   : list by id, example "id 2"
> save : save games list to json
> quit : exit program` + "\n")
		in := bufio.NewScanner(os.Stdin)
		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())

		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit", "q":
			return
		case "list", "l":
			fmt.Printf("%-8s %-15s %-10s %s\n", "ID", "Name", "Price", "Genre")
			for _, g := range games {
				fmt.Printf("#%-7d %-15s $%-9d %s\n", g.id, g.name, g.price, g.genre)
			}
		case "id":
			if len(cmd) < 2 {
				fmt.Println("Wrong ID")
				continue
			}

			num, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("ID should be a number")
				continue
			}

			g, ok := searchID[num]
			if !ok {
				fmt.Println("The game is not exist!")
				continue
			}
			fmt.Printf("%-8s %-15s %-18s %s\n", "ID", "Name", "Genre", "Price")
			fmt.Printf("#%-7d %-15s %-18s $%d\n", g.id, g.name, g.genre, g.price)
		case "save", "s":
			// load the data into the encodable struct
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable, jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}
			fmt.Println(string(out))
			return
		}
	}
}

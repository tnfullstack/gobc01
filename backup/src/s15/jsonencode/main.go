// jsonencode
package main

import (
	"encoding/json"
	"fmt"
)

// Permiss type
type permiss map[string]bool

type user struct {
	Name       string  `json:"username"`
	Passwd     string  `json:"-"`
	Permission permiss `json:",omitempty"`
}

func main() {

	users := []user{
		{Name: "inanc", Passwd: "1234", Permission: nil},
		{Name: "dog", Passwd: "42", Permission: permiss{"admin": true}},
		{Name: "devil", Passwd: "666", Permission: permiss{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(out))
}

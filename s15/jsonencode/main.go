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
	Passwd     string  `json:"***"`
	Permission permiss `json:"perms,omitempty"`
}

func main() {

	users := []user{
		{"inanc", "1234", nil},
		{"dog", "42", permiss{"admin": true}},
		{"devil", "666", permiss{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(out))
}

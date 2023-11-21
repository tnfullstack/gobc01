package main

import (
	"fmt"
	"os"
)

func main() {
	key := "DB_CONN"
	defValue := `postgres://as:name@example.com/pg? sslmode=verify-full`
	conn, err := SetEnv(key, defValue)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Set connection:", conn)
	}

	connStr := os.Getenv(key)
	fmt.Printf("Connection string: %s\n", connStr)

	if err = UnsetEnv(key); err != nil {
		fmt.Println(err)
	}

	conStr, err := GetenvDefault(key, defValue)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After unset connection:", conStr)
	}
}

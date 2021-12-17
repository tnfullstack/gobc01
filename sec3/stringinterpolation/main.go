package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	Name     string
	Age      int
	FaNum    float64
	OwnsADog bool
}

func main() {

	reader = bufio.NewReader(os.Stdin)
	var user User

	user.Name = stringReader("What is your name?")
	user.Age = readInt("How old are you?")
	user.FaNum = readFloat("What is your favourite Number?")
	user.OwnsADog = readBool("Do you own a dog?")

	fmt.Printf("You name is %s you are %d yeas old, your favourite number is %.2f. You own a dog: %t\n",
		user.Name, user.Age, user.FaNum, user.OwnsADog)
}

func prompt() {
	fmt.Print("-> ")
}

func stringReader(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("Please enter your name.")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println(userInput, "is not a number.")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println(userInput, "is not a number.")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		prompt()

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		} else {
			fmt.Println("Please enter 'Y' or 'N'.")
		}
	}
}

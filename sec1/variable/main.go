package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tnfullstack/gobc01/sec1/variable/doctor"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Can not read user input")
			return
		}
		userInput = strings.ToLower(userInput)

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" || userInput == "exit" {
			break
		}
		fmt.Println(doctor.Response(userInput))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"

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

		test := doctor.Response(userInput)

		fmt.Println(test)
	}
}

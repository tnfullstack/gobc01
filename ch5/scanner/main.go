package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func interview(str string) string {
	var returnStr string
	str = strings.ToLower(str)
	switch str {
	case "name":
		fmt.Printf("What is your %s? ", str)
		returnStr = getInfo()
	case "age":
		fmt.Printf("How is your %s? ", str)
		returnStr = getInfo()
	case "job":
		fmt.Printf("What is you %s? ", str)
		returnStr = getInfo()
	case "address":
		fmt.Printf("What is your %s? ", str)
		returnStr = getInfo()
	default:
		returnStr = fmt.Sprintf("%s is not a valid question.", str)
	}
	return returnStr
}

func getInfo() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func main() {
	fmt.Printf("%s is %s years old, he is a %s his address is %s\n", interview("name"),
		interview("age"), interview("job"), interview("address"))

	fmt.Println(interview("Bad"))
}

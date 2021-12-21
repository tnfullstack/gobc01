// infini loop
package main

import (
	"bufio"
	"fmt"
	"go/gobc01/sec5/infiniloop/logger"
	"os"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	ch := make(chan string)

	go logger.ListenForLog(ch)

	fmt.Println("Enter something.")

	for i := 1; i <= 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')

		ch <- input
		time.Sleep(time.Second)
	}

}

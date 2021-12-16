// Scope challenge
package main

import (
	"fmt"

	"github.com/tnfullstack/gobc01/sec2/scope/packageone"
)

var myVar = "Package level in main."

func main() {

	// Block level variable
	var blockVar = "Block level variable in main."

	fmt.Println(packageone.PrintMe(myVar), packageone.PrintMe(blockVar), packageone.PrintMe(packageone.PackageOneVar))

}

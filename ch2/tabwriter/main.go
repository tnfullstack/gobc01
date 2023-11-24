// Aligning text with tabwriter
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJohn\tSmitch\t")
	w.Flush()
}

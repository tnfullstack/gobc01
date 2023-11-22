// Reading/Writing from the child process
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var cmd string

	if runtime.GOOS == "windows" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	proc := exec.Command(cmd)

	buf := bytes.NewBuffer([]byte{})

	// The buffer which implements io.Writer interface is assigned
	// to stdout of the process
	proc.Stdout = buf

	// To avoid race conditions in this example. We wait till the
	// process exit.
	proc.Run()

	// The process writes the output to the buffer and we use the bytes
	// to print the output.
	fmt.Println(buf.String())
}

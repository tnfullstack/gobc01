package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	prcRun("ls", "-la")
	fmt.Println("-----------------------")
	prcStart("lsb", "-la")
}

func prcRun(cmd, flag string) {
	// defind an external command
	prc := exec.Command(cmd, flag)
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(prc.String(), out.String())
	}
}

func prcStart(cmd, flag string) {
	//defind an external command
	prc := exec.Command(cmd, flag)
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	if err := prc.Start(); err != nil {
		msg := "failed to start command %s %s"
		fmt.Println(msg, cmd, flag)
	}
	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(prc.String(), out.String())
	}
}

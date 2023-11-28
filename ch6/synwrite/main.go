// Writing the file from multiple goroutines
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type SyncWriter struct {
	m      sync.Mutex
	Writer io.Writer
}

func (w *SyncWriter) Write(b []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	return w.Writer.Write(b)
}

var data = []string{
	"Hello!",
	"Ola!",
	"Ahoj!",
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	wr := &SyncWriter{sync.Mutex{}, f}
	wg := sync.WaitGroup{}
	for _, val := range data {
		wg.Add(1)
		go func(greetings string) {
			fmt.Fprintln(wr, greetings)
			wg.Done()
		}(val)
	}
	wg.Wait()
}

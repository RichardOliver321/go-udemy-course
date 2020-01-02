package main

import (
	"fmt"
	"io"
	"os"
)

type terminalWriter struct{}

func main() {

	fn := os.Args[1]
	f, err := os.Open(fn)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	t := terminalWriter{}

	io.Copy(t, f)
}

func (terminalWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

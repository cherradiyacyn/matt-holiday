package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	n = len(p)
	*bc += ByteCounter(n)
	return n, nil
}

func main() {
	f1, _ := os.Open("in.txt")
	var bc ByteCounter
	f2 := &bc
	// look here, i used io.Copy for a different purpose; not to copy but just to count. it is made possible because io.Copy just needs a Writer, not specifically a file. that's an interesting thing about interfacesin Go.
	_, _ = io.Copy(f2, f1)
	fmt.Println("counted", bc, "bytes")
}

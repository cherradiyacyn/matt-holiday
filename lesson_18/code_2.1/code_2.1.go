package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f1, _ := os.Open("in.txt")
	f2, _ := os.Create("out.txt")
	n, _ := io.Copy(f2, f1)
	fmt.Println("copied", n, "bytes")
}

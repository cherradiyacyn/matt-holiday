package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	fmt.Println(hello.Greet(os.Args[1:]))
}

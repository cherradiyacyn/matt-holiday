package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(hello.Greet(os.Args[1]))
	} else {
		fmt.Println(hello.Greet("world"))
	}
}

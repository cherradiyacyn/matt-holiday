package main

import "fmt"

func main() {
	a, b := 2, 3.1
	fmt.Printf("a: %8T %[1]v\n", a)
	fmt.Printf("b: %8T %[1]v\n", b)

	a, b = int(b), float64(a)
	fmt.Printf("a: %8T %[1]v\n", a)
	fmt.Printf("b: %8T %[1]v\n", b)
}

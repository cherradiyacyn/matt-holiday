package main

import "fmt"

func do(b [3]int) int {
	b[0] = 0
	fmt.Printf("b@ %p\n", &b)
	return b[1]
}

func main() {
	a := [3]int{1, 2, 3}
	fmt.Printf("a@ %p\n", &a)
	v := do(a)
	fmt.Println(a, v)
}

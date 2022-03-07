package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

// Currying takes a function and reduces its arguments count by one
// (one argument gets bound, and a new function is returned)

func AddToA(a int) func(int) int {
	return func(b int) int {
		return Add(a, b)
	}
}

func main() {
	AddTo1 := AddToA(1)
	fmt.Println(AddTo1(2) == Add(1, 2))
}

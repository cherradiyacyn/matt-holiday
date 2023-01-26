package main

import "fmt"

func main() {
	var u = make([]int, 5)
	u[0] = 23 // it works!
	fmt.Println(u)
	var v = make([]int, 0, 5)
	v[0] = 23 // panic: runtime error: index out of range [0] with length 0
	fmt.Println(v)
}

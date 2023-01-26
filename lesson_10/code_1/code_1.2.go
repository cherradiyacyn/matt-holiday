package main

import "fmt"

func main() {
	var u = make([]int, 5)
	u = append(u, 23) // is NOT u[0] = 23
	fmt.Println(u)
	var v = make([]int, 0, 5)
	v = append(v, 23) // it works!
	fmt.Println(v)
}

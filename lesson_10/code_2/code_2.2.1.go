package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[:1]
	c := b[:2]
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	c = append(c, 4) // this will mutate the 3rd value of a :\
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
}

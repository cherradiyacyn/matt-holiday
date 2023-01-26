package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[:1]
	c := b[:2:2] // because of this :)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	c = append(c, 4) // now, it is safe :)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c) // notice that c is now in a different place %p :)
}

package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[:1]
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	c := b[:2]
	fmt.Println("c =", c)                            // WTF
	fmt.Printf("b: len %v cap %v\n", len(b), cap(b)) // this explains it!
	d := c[:1:1]                                     // [i:j:k]: len j-i cap k-i
	fmt.Printf("d: len %v cap %v\n", len(d), cap(d))
	e := d[:2] // panic: runtime error: slice bounds out of range [:2] with capacity 1
	fmt.Println(e)
}

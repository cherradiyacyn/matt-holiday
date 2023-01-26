package main

import "fmt"

func do(f func()) {
	f()
}

func main() {
	// i is in the same location on memory, even after 4 iterations, hence we'll get the same val from %p.
	for i := 0; i < 4; i++ {
		f := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
		do(f)
	}
}

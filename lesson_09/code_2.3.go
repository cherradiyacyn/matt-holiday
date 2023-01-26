package main

import "fmt"

func main() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		// after each iteration, create a new var and copy to it the val of i, so your closure will point to different places with different values in RAM.
		new_i := i
		s[i] = func() { fmt.Printf("%d @ %p\n", new_i, &new_i) }
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}

package main

import "fmt"

func main() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		// when i's enclosed every closure points to the same thing in memory : &i.
		s[i] = func() { fmt.Printf("%d @ %p\n", i, &i) }
	}

	// the scope is different so i from the first loop is NOT i in the second one.
	for i := 0; i < 4; i++ {
		s[i]()
		// Output :
		// 4 @ 0xc000018088
		// 4 @ 0xc000018088
		// 4 @ 0xc000018088
		// 4 @ 0xc000018088
		// Explanation :
		// the last val stored in &i is = 4, so the output is not a bug!
	}
}

package main

import "fmt"

func main() {
	s := "élite"
	// len(string) always returns the physical length in ram : 6
	// NOT the logical(visual) one : 5
	fmt.Printf("%T %[1]v %d\n", s, len(s))
	r := []rune(s)
	fmt.Printf("%T %[1]v %d\n", r, len(r))
	b := []byte(s)
	fmt.Printf("%T %[1]v %d\n", b, len(b))
}

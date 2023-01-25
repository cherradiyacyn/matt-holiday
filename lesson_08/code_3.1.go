package main

import "fmt"

// we pass a copy of a descriptor so m and a are two distinct maps pointing to the same hashtable.
func do(m map[string]int) {
	m["d"] = 3
	// here m will point to a new hashtable.
	m = make(map[string]int)
	m["e"] = 23
	fmt.Println("m", m)
}

func main() {
	a := map[string]int{"a": 2, "b": 0, "c": 2}
	fmt.Println("a", a)
	do(a)
	fmt.Println("a", a)

}

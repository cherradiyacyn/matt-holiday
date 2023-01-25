package main

import "fmt"

// we pass a pointer to a descriptor so m and a are now the same thing.
func do(m *map[string]int) {
	(*m)["d"] = 3
	// if m is changed, a is changed.
	(*m) = make(map[string]int)
	(*m)["e"] = 23
	fmt.Println("m", *m)
}

func main() {
	a := map[string]int{"a": 2, "b": 0, "c": 2}
	fmt.Println("a", a)
	do(&a)
	fmt.Println("a", a)

}

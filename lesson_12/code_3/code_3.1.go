package main

import "fmt"

type struct1 struct {
	field1 string
	field2 int
}

type struct2 struct {
	field1 string
	field2 int
}

func main() {
	var a = struct1{"abc", 1}
	fmt.Println(a)
	var b = struct2{"def", 2}
	fmt.Println(b)
	// OK; same names & types & order
	a = struct1(b)
	fmt.Println(a, b)
}

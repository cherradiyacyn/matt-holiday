package main

import "fmt"

type album1 struct {
	title string
}

type album2 struct {
	title string
}

func main() {
	var a1 = album1{
		"The white album",
	}
	var a2 = album2{
		"The black album",
	}
	a1 = album1(a2)
	fmt.Println(a1, a2)
}

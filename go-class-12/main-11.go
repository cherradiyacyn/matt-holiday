package main

import "fmt"

func main() {
	var album1 = struct {
		title string
	}{
		"The white album",
	}
	var album2 = struct {
		title string
	}{
		"The black album",
	}
	album1 = album2
	fmt.Println(album1, album2)
}

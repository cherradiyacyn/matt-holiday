package main

import "fmt"

func main() {
	var album = struct {
		title string
	}{
		"The white album",
	}

	var ptrAlbum = &album
	var ptrAlbum2 *struct {
		title string
	}

	fmt.Println(album, ptrAlbum, ptrAlbum2)
}

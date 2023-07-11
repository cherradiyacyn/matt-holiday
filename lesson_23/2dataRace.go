package main

import (
	"fmt"
	"net/http"
)

var nextID int

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div>You got %d</div>", nextID)
	nextID++ // UNSAFE
}

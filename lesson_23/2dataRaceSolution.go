package main

import (
	"fmt"
	"net/http"
)

var nextID = make(chan int)

func main() {
	go counter()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div>You got %d</div>", <-nextID)
}

func counter() {
	for i := 0; ; i++ {
		nextID <- i
	}
}

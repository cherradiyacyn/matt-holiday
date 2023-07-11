package main

import (
	"fmt"
	"net/http"
)

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div>You got %d</div>", <-ch)
}

func main() {
	var nextID nextCh = make(chan int)
	go counter(nextID)
	http.HandleFunc("/", nextID.handler)
	http.ListenAndServe(":8080", nil)
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

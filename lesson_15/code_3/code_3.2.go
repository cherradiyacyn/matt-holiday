package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todoTemplate = `
<h1>Todo #{{.Id}}</h1>
<div>{{printf "User %d" .UserId}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {
	const url = "https://jsonplaceholder.typicode.com/"

	resp, err := http.Get(url + r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	var item todo

	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	template := template.New("todo")
	template.Parse(todoTemplate)
	template.Execute(w, item)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

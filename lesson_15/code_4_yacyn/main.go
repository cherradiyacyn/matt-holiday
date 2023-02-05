package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type Quote struct {
	Id        int    `json:"id"`
	QuoteText string `json:"quote"`
	Author    string `json:"author"`
}

var quoteTemplate = `
<h1>Quote #{{.Id}}</h1>
<div>{{.QuoteText}}</div>
<div>{{printf "—%s" .Author}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {
	const url = "https://dummyjson.com/"
	resp, err := http.Get(url + r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	var item Quote
	if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	template := template.New("quote")
	template.Parse(quoteTemplate)
	template.Execute(w, item)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

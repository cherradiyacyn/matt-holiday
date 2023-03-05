package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32 // BAD, but this is just h/w.

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// add the handlers.
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q\n", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q\n", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])
}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "item: %q has the price: %q\n", item, db[item])
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q\n", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "new price: %q for item: %q updated succesfully\n", db[item], item)
}

func (db database) drop(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "item: %q deleted\n", item)
}

func main() {
	// define the database.
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	// add some routes.
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/read", db.fetch)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.drop)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

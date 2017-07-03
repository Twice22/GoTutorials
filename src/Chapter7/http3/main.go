package main 

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := databse{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

// handler for /list
// db.list is a method value (see 6.4), that is to say, a value of type
// func(w http.ResponseWriter, req *http.Request)
// when called, invokes the database.list method with receiver db
// http.HandlerFunc is a conversion
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// handler for price
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such items: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
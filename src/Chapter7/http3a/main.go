package main 

import (
	"fmt"
	"log"
	"net/http"
)

/*  HandlerFunc satisfies
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
*/

func main() {
	db := database{"shoes": 50, "socks": 5}

	// ServeMux has a convenience method called HandleFunc
	// that allows us to simplify the handler registration code.
	mux := http.NewServeMux()

	// equivalent to
	// mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8001", mux))
}

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: $%d\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "$%d\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
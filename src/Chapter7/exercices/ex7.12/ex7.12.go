package main 

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strconv"
	"sync"
)

/*  HandlerFunc satisfies
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
*/

var dbmux sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()

	// net/http provides a global ServeMUx instance called DefaultServeMux
	// and package-level fcts called http.Handle and http.HandleFunc.
	mux.HandleFunc("/list", db.list) // use to read
	mux.HandleFunc("/price", db.price) // use to read
	mux.HandleFunc("/create", db.create) // use to create or update an item
	mux.HandleFunc("/delete", db.delete) // use to delete an item

	log.Fatal(http.ListenAndServe("localhost:8001", mux)) // pass mux (and not nil)

}

type database map[string]int


var tableLIst = template.Must(template.New("tableList").Parse(`
		<h1>List</h1>
		<table>
		<tr style='text-align: left'>
			<th>Item</th>
			<th>Price</th>
		</tr>
		{{ range $key, $value := .}}
		<tr>
			<td>{{$key}}</td>
			<td>${{$value}}</td>
		</tr>
		{{end}}
		</table>
	`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	dbmux.Lock()
	if err := tableLIst.Execute(w, &db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to execute template: %q\n", err)
	}
	dbmux.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	dbmux.Lock()
	price, ok := db[item]
	dbmux.Unlock()
	if ok {
		fmt.Fprintf(w, "$%d\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price_str := req.URL.Query().Get("price")

	price, err := strconv.Atoi(price_str)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong price: %s\n", price_str)	
	} else {
		dbmux.Lock()
		db[item] = price
		dbmux.Unlock()
		fmt.Fprintf(w, "created or updated %s: $%d\n", item, price)
	}
}


func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		dbmux.Lock()
		delete(db, item)
		dbmux.Unlock()
		fmt.Fprintf(w, "delete item: %q\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}	
}
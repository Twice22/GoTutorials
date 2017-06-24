package main 

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Use Mutex to avoid race condition to happen
// mu.Lock() and mu.Unlock() ensures that at most one goroutine accesses
// the variable at a time.
var mu sync.Mutex
var count int

func main() {
	// any other requests beside /count invoke handler
	http.HandleFunc("/", handler)

	// /count invokes counter
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
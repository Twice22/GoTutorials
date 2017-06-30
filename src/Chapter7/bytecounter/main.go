package main 

import (
	"fmt"
)

type ByteCounter int

// Write method of *ByteCounter type
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // Convert int to ByteCounter (not a fct call)
	return len(p), nil
}

func main() {
	var c ByteCounter
	(&c).Write([]byte("hello")) // can use c.Write also
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"

	// Fprintf(w io.Writer, format string, arg ...interface{}) (int, error)
	// and *ByteCounter satisfies the io.Writer contract because it defines a Write method
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}

/* io.Write contract

package io

type Writer interface {
	Write(p []byte) (n int, err error)
}


*/
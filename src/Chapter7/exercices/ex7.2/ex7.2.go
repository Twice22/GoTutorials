package main 

import (
	"fmt"
	"io"
	"os"
)

type byteCounter struct {
	nw io.Writer
	sz int64
}

func (c *byteCounter) Write(p []byte) (int, error) {
	sz, err := c.nw.Write(p) // call Write method from io.Writer interface
	c.sz += int64(sz)
	return sz, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	nw := byteCounter{w, 0}
	return &nw, &nw.sz
}

func main() {
	writer, sz := CountingWriter(os.Stdout)

	// call Write from *byteCounter type and populate the size sz
	fmt.Fprint(writer, "Hello world\n")

	// print the size of length of the text in writer
	fmt.Println(*sz)
}
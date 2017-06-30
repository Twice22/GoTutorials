package main 

import (
	"fmt"
	"io"
	"strings"
	"bytes"
)


type LimitedReader struct {
	lr io.Reader
	limit int
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	// read at max "limit" characters from r
	n, err = r.lr.Read(p[: r.limit])
	if n >= r.limit {
		err = io.EOF
	}
	// bare return
	return
}

func LimitReader(r io.Reader, n int) io.Reader {
	// we can return a LimitedReader as a io.Reader
	// as io.Reader is embedded in LimitedReader
	return &LimitedReader{r, n}
}


func main() {
	r := LimitReader(strings.NewReader("I want to read only 10 characters"), 10)

	b := new(bytes.Buffer)

	// r need to satisfy io.Reader interface (it is the case as io.Reader is embedded by LimitedReader)
	nb, _ := b.ReadFrom(r)

	if nb == 10 {
		fmt.Println(b.String())
	}

}
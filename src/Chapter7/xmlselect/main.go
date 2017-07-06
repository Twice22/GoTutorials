package main 

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// go build Chapter1/Fetch
// ./Fetch.exe https://www.w3.org/TR/2006/REC-xml11-20060816 | go run main.go div div h2
func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		// type switch. switch if type (inside ())
		// if one of the case type (and use method expression accordingly)
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// reports wether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
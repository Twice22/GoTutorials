package main 

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type StringReader struct {
	s string
}

// implement implement io.Reader interface
func(r *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, []byte(r.s))
	err = io.EOF // set err to EOF to tell to the Parser that we reached end of file
	// bare return
	return
}

func NewReader(str string) *StringReader {
	s := StringReader{str}
	return &s
}


func main() {
	doc, _ := html.Parse(NewReader(`
		<html>
			<head>
				<title>My WebSite</title>
			</head>
			<body>
			<div>
				<p>This is a paragraph</p>
			</div>
			</body>
		</html>`))

	fmt.Println(doc.FirstChild.FirstChild.Data) // "head"

}
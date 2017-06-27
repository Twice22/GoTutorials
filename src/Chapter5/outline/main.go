package main 

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// prints the outline of an HTML document tree
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// go build Chapter1/fetch
// go get golang.org/x/net/html
// go build Chapter5/outline
// ./fetch.exe https://golang.org | ./outline.exe
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
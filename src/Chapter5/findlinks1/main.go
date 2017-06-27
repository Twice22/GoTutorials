package main 

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)


// visit appends to links each link found in n and returns the result
func visit(links []string, n *html.Node) []string {
	// if <a>
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			// add value of href attribute to links
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// recursively call visit for each of n's children held in FirstChild linked list
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// go build Chapter1/fetch
// go get golang.org/x/net/html
// go build Chapter5/findlinks1
// ./fetch.exe https://golang.org | ./findlinks1.exe
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}


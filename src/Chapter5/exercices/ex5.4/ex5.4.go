package main 

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)


// visit appends to links each link found in n and returns the result
func visit(links []string, n *html.Node) []string {
	// if <a>
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "src" {
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
// go build Chapter5/exercices/ex5.4
// ./fetch.exe https://golang.org | ./ex5.4.exe
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}


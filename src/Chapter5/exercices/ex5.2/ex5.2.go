package main 

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)


// visit appends to links each link found in n and returns the result
func counts(tagCounts map[string]int, n *html.Node) map[string]int {
	// if <a>
	if n.Type == html.ElementNode {
		tagCounts[n.Data]++
	}
	// recursively call visit for each of n's children held in FirstChild linked list
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tagCounts = counts(tagCounts, c)
	}
	return tagCounts
}

// go build Chapter1/fetch
// go get golang.org/x/net/html
// go build Chapter5/exercices/ex5.2
// ./fetch.exe https://golang.org | ./ex5.2.exe
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for tag, n := range counts(make(map[string]int), doc) {
		fmt.Printf("%s\t%d\n", tag, n)
	}
}


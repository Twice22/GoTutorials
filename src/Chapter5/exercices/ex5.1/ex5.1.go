package main 

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)


func visit(links []string, n *html.Node) []string {
	if n != nil {
		links = visit(links, n.FirstChild)

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				// add value of href attribute to links
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}

		links = visit(links, n.NextSibling)
	}
	return links
}

// go build Chapter1/fetch
// go get golang.org/x/net/html
// go build Chapter5/exercices/ex5.1
// ./fetch.exe https://golang.org | ./ex5.1.exe
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


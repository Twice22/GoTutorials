package main 

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// * adverb in %*s prints a string padded
		// with depth*2 width of string ""
		fmt.Printf("%*s<%s", depth*2, "", n.Data)

		// print Attribute and value if any
		for _, a := range n.Attr {
			// use %q to account for the quotes if any
			fmt.Printf(" %s=%q", a.Key, a.Val)
		}
		end := ">\n"
		if n.FirstChild == nil {
			end = "/>\n"
		}
		fmt.Printf("%s", end)

		depth++
	} else if n.Type == html.TextNode {
		trimStr := strings.TrimSpace(n.Data)
		if trimStr != "" {
			fmt.Printf("%*s%s\n", depth*2, "", trimStr)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}		
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// go run ex5.7.go http://golang.org
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}
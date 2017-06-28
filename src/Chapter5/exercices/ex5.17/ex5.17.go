package main 

import (
	"fmt"
	"net/http"
	"os"
	"bytes"

	"golang.org/x/net/html"
)

func contains(s string, slice []string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}


func ElementsByTagName(doc *html.Node, name ...string) (res []*html.Node) {
	startElement := func(n *html.Node) bool {
		if n.Type == html.ElementNode && contains(n.Data, name) {
			return false
		}
		return true
	}

	return forEachNode(doc, startElement, nil, nil)
}

// change forEachNode to return the corresponding node
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool, find []*html.Node) []*html.Node {
	if pre != nil {
		if !pre(n) {
			find = append(find, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		find = forEachNode(c, pre, post, find)
	}

	if post != nil {
		if !post(n) {
			find = append(find, n)
		}
	}
	// bare return (return res)
	return find
}

func printTag(n *html.Node) string {
	var buf bytes.Buffer

	if n.Type == html.ElementNode {
		fmt.Fprintf(&buf, "<%s", n.Data)
		for _, a := range n.Attr {
			// use %q to account for the quotes if any
			fmt.Fprintf(&buf, " %s=%q", a.Key, a.Val)
		}
		end := ">\n"
		if n.FirstChild == nil {
			end = "/>\n"
		}
		buf.WriteString(end)
		return buf.String()
	}
	return ""
}

func printTags(n []*html.Node) string {
	var buf bytes.Buffer

	for _, elt := range n {
		buf.WriteString(printTag(elt))
		buf.WriteString("\n")
	}
	return buf.String()
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

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	// helper function to ensure we've found the
	// right element
	fmt.Printf(printTags(images))
	fmt.Printf(printTags(headings))

	return nil
}

// go run ex5.17.go http://koreus.com
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}
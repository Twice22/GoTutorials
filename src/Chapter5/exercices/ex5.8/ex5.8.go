package main 

import (
	"fmt"
	"net/http"
	"os"
	"bytes"

	"golang.org/x/net/html"
)


func ElementByID(doc *html.Node, id string) *html.Node {

	// define startElement inside ElementByID in order to access
	// id string variable (otherwise we would need to pass id as a
	// parameter to startElement)
	startElement := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					return false
				}
			}
		}
		return true // if true keep searching
	}

	return forEachNode(doc, startElement, nil)
}

// change forEachNode to return the corresponding node
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		if find := forEachNode(c, pre, post); find != nil {
			return find
		}
	}

	if post != nil {
		if !post(n) {
			return n
		}
	}
	return nil
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

	doc = ElementByID(doc, "footer")

	// helper function to ensure we've found the
	// right element
	fmt.Printf(printTag(doc))

	return nil
}

// go run ex5.8.go http://golang.org
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}
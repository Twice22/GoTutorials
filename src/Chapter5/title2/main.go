package main 

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n * html.Node)) {
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

// fetch an HTML document and prints its title
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// defer call. The function and argument expression if any are
	// evaluated when the statement is executed, but the actual call
	// is deferred until the function that contains the defer statement
	// has finished (here when title function has finished)
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
				fmt.Println(n.FirstChild.Data)
			}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}
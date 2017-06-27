package main 

import (
	"fmt"
	"os"
	"net/http"
	"bufio"
	"strings"
	"golang.org/x/net/html"
)

func countWordsAndImages(n *html.Node) (words, images int) {
	// images count
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))

		// set the split function to split words
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	// bare return
	return 
}


func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		// bare return (can be used when return variables are named)
		// it is equivalent to:
		// return words, images, err
		// if  return variables are not populated then it will return the zero
		// values for these variables
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		// bare return
		return
	}
	// call countWordsAndImages (lower c in count!)
	words, images = countWordsAndImages(doc)
	return
}

// go run ex5.5.go http://golang.org
// go run ex5.5.go http://golang.org https://www.koreus.com
func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex5.5: %v\n", err)
			continue
		}
		fmt.Printf("%s\t%d images\t%d words\n", url, images, words)
	}
}

package main 

import (
	"fmt"
	"log"
	"os"

	"Chapter5/links"
)

// breadFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		// at first items = ["https://golang.org"] for example
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// we append all the url found in "https://golang.org" page
				// and add it to worklist, then we loop. Now items = ['link1', 'link2',...]
				// and we append all the links we find in each linkX pages
				// seen[item] avoid to loop indefinetely (indefinetely add the same URLs)
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}


// go run findlinks https://golang.org
func main() {
	breadFirst(crawl, os.Args[1:])
}
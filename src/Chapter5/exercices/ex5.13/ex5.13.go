package main 

import (
	"fmt"
	"log"
	"os"
	"io"
	"strings"
	"net/url"
	"path/filepath"
	"net/http"

	"Chapter5/links"
)

// host hold the original host
// it is populated with the original host in main
var originHost string

func save(myUrl string) (err error) {
	u, err := url.Parse(myUrl)
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(originHost, u.Host) {
		return nil
	}
	dir := filepath.Join(u.Host, u.Path)
	filename := filepath.Join(dir, "index.html")

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	resp, err := http.Get(myUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", myUrl, resp.Status)
	}

	// create file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	// copy resp.Body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	
	err = file.Close()
	if err != nil {
		return err
	}
	return nil

}

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
	err := save(url)
	if err != nil {
		log.Printf("can't save %q: %s", url, err)
	}

	list, err := links.Extract(url)
	if err != nil {
		log.Printf("can't extract links %q: %s", url, err)
	}
	return list
}


// go run findlinks https://golang.org
func main() {
	originHost = os.Args[1]
	breadFirst(crawl, os.Args[1:])
}
package main 

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	// create channel of strings
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go statement creates additional goroutine
		// that call fetch asynchronously
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		// Note: when one goroutine attempts a send or receive on a channel,
		// it blocks until another goroutine attempts the corresponding receive or send operation
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// read the body of the response and discard it by writing
	// to the ioutil.Discard output stream
	// Copy returns the bytes count and err
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak ressources
	if err != nil {
		// send a summary line on the channel ch
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
package main 

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

// it will try to connect every: 1s, 2s, 4s, 8s, 16s.. until 1 minute
// go run main.go http://google.xcom
// (Note: xcom instead of com, so we are sure the url doesn't exist!)
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// log.Fatalf("Site is down: %v\n", err)
		// log.Fatalf will prefixe the error with time and date

		// We can also set the prefix used by the log package to the
		// name of the command, and suppress the date and time using:
		// log.setPrefix("wait: ")
		// log.SetFlags(0)
		os.Exit(1)
	}
}
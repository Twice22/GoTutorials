package main 

import (
	"fmt"
	"flag"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

// go run main.go
// go run main.go -period 50ms
// go run main.go -period 1m15
// go run main.go -period "1 day"
func main() {
	flag.Parse()

	// the fmt package calls the time.Duration's String
	// method to print the period in a user-friendly notation
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
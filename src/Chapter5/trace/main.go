package main 

import (
	"log"
	"time"
)

func bigSlowOperation() {
	f := trace("bigSlowOperation")
	defer f()

	// we can also write the 2 lines above as one:
	//defer trace("bigSlowOperation")()

	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

// function that returns a function
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// we can see that f is executed after 10 seconds due to the defer statement
func main() {
	bigSlowOperation()
}
package main 

import (
	"fmt"
	"time"
)

func main() {
	// go f() // create a new goroutine that calls f(); don't wait for f() to
	// return before executing the following code.
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow because we are not using memoization
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			// \r allows to return to the beginning of the line
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
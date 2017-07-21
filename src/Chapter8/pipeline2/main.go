package main 

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x <100 ; x++ {
			// channel are unbuffered so, x is added to the channel
			// then the channel is "full" so we need to remove the element
			// from the channel (the x := <- naturals from the second goroutine)
			// is then called
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	// we can use range over a channel to retrieve one value at a time
	// until the channel is closed (using close fct)
	for x := range squares {
		fmt.Println(x)
	}
}

// Note: We can also retrieve the element and a boolean from a channel using:
/*
for {
	x, ok := <- naturals
	if !ok {
		break
	}
	squares <- x * x
}
close(squares)
*/
// if ok is false then we couldn't retrieve the value. That means the channel was closed and drained
// and so we use a break statement to stop looping
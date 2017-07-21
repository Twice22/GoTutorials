package main 

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			// channel are unbuffered so, x is added to the channel
			// then the channel is "full" so we need to remove the element
			// from the channel (the x := <- naturals from the second goroutine)
			// is then called
			naturals <- x
		}
	}()

	go func() {
		for {
			// once the channel is full (one element in the channel because unbuffered chan)
			// retrieve this element and pass the square of this element to the squares channel
			x := <- naturals

			// if there is already one element in the squares channel, we need to remove int. To do
			// so <-squares is called from the main goroutine below
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
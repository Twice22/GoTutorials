package main 

import "fmt"

// chan<- is a sending-only channel
// <-chan is a receiving-only channel
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}	
}

func main() {
	// naturals and squares are both use as receiving and sending channels
	// so we need to create bidirectional channels
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
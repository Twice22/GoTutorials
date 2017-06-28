package main 

import "fmt"

func recov() (res int) {
	defer func() {
		recover()
		res = 2
	}()

	// we panic, so the function will execute
	// the deferred function that will recover
	// from the panic (recover()) and assign 2 to res
	panic(8)
}

func main() {
	fmt.Printf("%d\n", recov())
}
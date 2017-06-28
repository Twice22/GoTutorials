package main

import "fmt"

// squares returns a function that returns
// the nest square number each time it is called

// Note: square return a type "func() int"
// that is why we have "return func() int {}"
// inside the fct
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Printf("%T\n", f) // func() int
	// as f is type func() int, taking f() will return an int (here x*x)
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
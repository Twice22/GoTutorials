package main 

import "fmt"


// reverse a slice of ints in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// using "..." define an array (not a slice!) whose
	// length is computed from the number of initial elements
	a := [...]int{0, 1, 2, 3, 4, 5}

	reverse(a[:]) // we pass a slice using [:]
	fmt.Println(a)

	// define a slice
	s := []int{0, 1, 2, 3, 4, 5}

	// Rotate s left by 2 positions
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
}
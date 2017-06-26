package main 

import "fmt"


// reverse a slice of ints in place
// we must provide the size if we use a pointer on array
// that is why pointer on array are only used when we know
// the size (for cryptographic fct such as SHA256 for example)
func reverse(s *[6]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// using "..." define an array (not a slice!) whose
	// length is computed from the number of initial elements
	a := [...]int{0, 1, 2, 3, 4, 5}

	reverse(&a) // we pass a slice using [:]
	fmt.Println(a)
}
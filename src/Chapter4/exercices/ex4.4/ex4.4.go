package main 

import "fmt"


func rotate(s []int, r int) {
	length := len(s)
	r = r % length

	// copy r first element in temp
	temp := make([]int, r)
	copy(temp, s[:r])

	// add element from r to end of s to beginning of s
	copy(s, s[r:])

	// append temp after previous copied elements
	copy(s[length-r:], temp)
}

func main() {
	// using "..." define an array (not a slice!) whose
	// length is computed from the number of initial elements
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	rotate(a[:], 22) // we pass a slice using [:]
	fmt.Println(a)
}
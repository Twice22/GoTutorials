package main 

import "fmt"


func removeDup(s []string) []string {
	i := 0

	for {
		if i+1 >= len(s) {
			break
		}
		if s[i+1] == s[i] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}

func main() {
	// using "..." define an array (not a slice!) whose
	// length is computed from the number of initial elements
	a := [...]string{"dup", "adj", "adj", "adj", "non", "adj", "dup", "dup", "tit", "dup", "dup", "dup"}

	fmt.Println(removeDup(a[:]))
}
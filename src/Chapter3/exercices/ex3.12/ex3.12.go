package main 

import (
	"fmt"
	"sort"
	"strings"
	"os"
)

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// convert string to slice of strings
	b1 := strings.Split(s1, "")
	b2 := strings.Split(s2, "")

	// sort string slices
	sort.Strings(b1)
	sort.Strings(b2)

	// convert the sorted slice back to a string
	s1 = strings.Join(b1, "")
	s2 = strings.Join(b2, "")

	return s1 == s2
}

func main() {
	if len(os.Args) >= 3 {
		fmt.Println("anagram: ", anagram(os.Args[1], os.Args[2]))
	}
}
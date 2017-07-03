package main 

import (
	"fmt"
	"sort"
)

func IsPalindrom(s sort.Interface) bool {
	length := s.Len()
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

type Palindrom struct {
	t []byte
}

func (x Palindrom) Len() int { return len(x.t) }
func (x Palindrom) Less(i, j int) bool { return x.t[i] < x.t[j] }
func (x Palindrom) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	fmt.Println(IsPalindrom(Palindrom{[]byte("radar")}))
	fmt.Println(IsPalindrom(Palindrom{[]byte("checking")}))
	fmt.Println(IsPalindrom(Palindrom{[]byte("IuytrrtyuI")}))
}
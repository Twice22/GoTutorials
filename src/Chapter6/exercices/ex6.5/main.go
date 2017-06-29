package main

import (
	"fmt"
	"Chapter6/exercices/ex6.5/intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	for _, e := range x.Elems() {
		fmt.Printf("%d ", e)
	}

}
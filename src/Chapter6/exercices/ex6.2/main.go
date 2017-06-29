package main

import (
	"fmt"
	"Chapter6/exercices/ex6.2/intset"
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

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Len()) // "4"

	x.Remove(42)

	fmt.Println(&x) // "{1 9 144}"

	//x.Clear()

	//fmt.Println(&x)

	// 1 already exist in the set so it won't be added!
	x.AddAll(1, 888, 47, 92)

	fmt.Println(&x)

}
package main

import (
	"fmt"
	"Chapter6/exercices/ex6.1/intset"
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

	x.Clear()

	fmt.Println(&x)

	x.Add(1)
	x.Add(8)
	x.Add(47)

	fmt.Println(&x)

	v := x.Copy()

	v.Remove(8)

	fmt.Println(&x)
	fmt.Println(v)

}
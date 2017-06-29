package main

import (
	"fmt"
	"Chapter6/exercices/ex6.3/intset"
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

	// Uncomment lines below to test Intersection
	//x.IntersectWith(&y)
	//y.IntersectWith(&x)
	//fmt.Println(&y)

	// Uncomment lines below to test Difference
	//x.DifferenceWith(&y)
	//fmt.Println(&x)

	// Uncomment lines below to test SymetricDifference
	x.SymetricDifference(&y)
	fmt.Println(&x)

}
package main

import (
	"fmt"
	"math/rand"
	"Chapter7/exercices/ex7.3/treesort"
)

func main() {
	data := make([]int, 50)
	var t *treesort.Tree

	
	for i := range data {
		data[i] = rand.Int() % 50
		// add number to the tree
		t = treesort.Add(t, data[i])
	}

	fmt.Println("Numbers in the tree (sorted method):")
	treesort.Sort(data)
	fmt.Printf("%v\n", data)
	
	fmt.Println("\nNumbers in the tree (string method):")
	fmt.Println(t)
}
package main

import (
	"Chapter2/popcount/popcount"
	"fmt"
)

func main() {
	var a uint64 = 2778
	var nbBytes int
	nbBytes = popcount.PopCount(a)
	fmt.Printf("Number of set bytes: %d\n", nbBytes)
}
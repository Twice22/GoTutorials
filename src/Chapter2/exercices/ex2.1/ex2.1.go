package main

import (
	"fmt"
	"Chapter2/exercices/ex2.1/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
	fmt.Println(tempconv.CToK(tempconv.BoilingC))
	fmt.Println(tempconv.CToF(37.5))
	fmt.Println(tempconv.FToC(42.0))
	fmt.Println(tempconv.FToK(42.0))
}